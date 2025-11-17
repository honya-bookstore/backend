local basename = vim.fs.basename
local env = vim.env
local g = vim.g
local nc_available = vim.fn.executable("nc")
local psql_available = vim.fn.executable("psql")

local is_atlasgo_community = true

local db_username = env.DB_USERNAME
local db_password = env.DB_PASSWORD
local db_host = env.DB_HOST
local db_port = env.DB_PORT
local db_database = env.DB_DATABASE

local psql_cmd = string.format(
  'PGPASSWORD="%s" psql -h %s -p %s -U %s -d %s -c "\\q" >/dev/null 2>&1',
  db_password,
  db_host,
  db_port,
  db_username,
  db_database
)

---@return boolean
function _G._IsDbUp()
  if nc_available == 1 then
    local nc_cmd = string.format("nc -z %s %s", db_host, db_port)
    local nc_status = os.execute(nc_cmd)
    if nc_status == 0 then
      return true
    end
  end

  if psql_available == 1 then
    local psql_status = os.execute(psql_cmd)
    if psql_status == 0 then
      return true
    end
  end

  return false
end

---@module 'lazy'
---@type LazySpec
return {
  {
    "stevearc/conform.nvim",
    ---@module 'conform'
    ---@type conform.setupOpts
    opts = {
      formatters = {
        swag_fmt = {
          command = "swag",
          args = {
            "fmt",
            "-d",
            "$FILENAME",
          },
          condition = function(_, ctx)
            local filename = ctx.filename
            return filename:match("internal/interface/api/handler") ~= nil
              or filename:match("internal/interface/api/mapper") ~= nil
              or filename:match("internal/interface/api/request") ~= nil
              or filename:match("internal/interface/api/response") ~= nil
              or filename:match("cmd") ~= nil
          end,
          stdin = false,
        },
        swag_gen = {
          command = "swag",
          args = {
            "init",
            "-g",
            "./cmd/main.go",
            "-ot",
            "go",
          },
          condition = function(_, ctx)
            local filename = ctx.filename
            return not g.dev_no_gen
              and (
                filename:match("internal/interface/api/handler") ~= nil
                or filename:match("internal/interface/api/mapper") ~= nil
                or filename:match("internal/interface/api/request") ~= nil
                or filename:match("internal/interface/api/response") ~= nil
                or filename:match("cmd") ~= nil
              )
          end,
          stdin = false,
        },
        wire = {
          command = "wire",
          args = {
            "gen",
            "$DIRNAME",
          },
          condition = function(_, ctx)
            return not g.dev_no_gen and basename(ctx.filename):match(".*%.wire.go") ~= nil
          end,
          stdin = false,
        },
        sqlc = {
          command = "sqlc",
          args = {
            "generate",
          },
          condition = function(_, ctx)
            local filename = ctx.filename
            return not g.dev_no_gen
              and (
                filename:match("database/.*schema%.sql") ~= nil
                or filename:match("database/queries/.*%.sql") ~= nil
                or filename:match("sqlc.yaml") ~= nil
              )
          end,
        },
        atlas = {
          command = "atlas",
          args = {
            "schema",
            "apply",
            "--env=local",
            "--auto-approve",
          },
          stdin = false,
          condition = function(_, ctx)
            if not _IsDbUp() or g.dev_no_gen then
              return false
            end
            local filename = ctx.filename
            if is_atlasgo_community then
              return filename:match("database/.*schema%.sql") ~= nil
            end
            return filename:match("database/.*%.sql") ~= nil and filename:match("database/.*seed-fake%.sql") == nil
          end,
        },
        pgsqlfluff = {
          condition = function(_, ctx)
            local filename = ctx.filename
            return filename:match("database/queries/.*%.sql") == nil
          end,
        },
        mockery = {
          command = "mockery",
          condition = function(_, ctx)
            local filename = ctx.filename
            return not g.dev_no_gen and filename:match(".*/internal/domain/%w*/repository%.go") ~= nil
          end,
        },
      },
      formatters_by_ft = {
        go = {
          "swag_fmt",
          "swag_gen",
          "wire",
          "mockery",
        },
        yaml = {
          "sqlc",
        },
        pgsql = {
          "atlas",
          "sqlc",
        },
      },
    },
    opts_extend = {
      "formatters_by_ft.go",
      "formatters_by_ft.yaml",
      "formatters_by_ft.pgsql",
    },
    optional = true,
  },
  {
    "mfussenegger/nvim-lint",
    opts = function()
      local lint = require("lint")

      if lint.linters.pgsqlfluff then
        local pgsqlfluff = lint.linters.pgsqlfluff
        lint.linters.pgsqlfluff = function()
          local bufname = vim.api.nvim_buf_get_name(0)
          if bufname:match("database/queries/.*%.sql") ~= nil then
            return {}
          end
          if type(pgsqlfluff) == "table" then
            return pgsqlfluff
          else
            return pgsqlfluff()
          end
        end
      end

      -- lint.linters.sqlc = function()
      --   local bufname = vim.api.nvim_buf_get_name(0)
      --   if bufname:match("database/(schema|queries/.*)%.sql") and _IsDbUp() then
      --     ---@type lint.Linter
      --     return {
      --       name = "sqlc",
      --       cmd = "sqlc",
      --       args = { "vet" },
      --       stream = "stderr",
      --       parser = require("lint.parser").from_pattern(
      --         "^(.+): (.+: .+): (.+)$",
      --         { "file", "code", "message" },
      --         nil,
      --         {
      --           source = "sqlc",
      --           severity = vim.diagnostic.severity.WARN,
      --         }
      --       ),
      --     }
      --   end
      --   return {}
      -- end

      lint.linters_by_ft.pgsql = lint.linters_by_ft.pgsql or {}
      -- table.insert(lint.linters_by_ft.pgsql, "sqlc")
    end,
    optional = true,
  },
  {
    "kristijanhusak/vim-dadbod-ui",
    opts = function()
      vim.g.dbs = {
        {
          name = "db-local",
          url = string.format(
            "postgres://%s:%s@%s:%s/%s?sslmode=disable",
            db_username,
            db_password,
            db_host,
            db_port,
            db_database
          ),
        },
        {
          name = "redis-local",
          url = "redis:0",
        },
      }
    end,
    optional = true,
  },
  {
    "folke/which-key.nvim",
    ---@module 'which-key'
    ---@type wk.Opts
    opts = {
      spec = {
        { "<localleader>l", group = "LSP", icon = { icon = "", color = "yellow" } },
      },
    },
    opts_extend = {
      "spec",
    },
    optional = true,
  },
}
