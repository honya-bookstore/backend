package repositorypostgres

import (
	"errors"

	"backend/internal/domain"

	"github.com/hashicorp/go-multierror"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func toDomainError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return multierror.Append(domain.ErrNotFound, errors.New("record not found"), err)
	}

	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) {
		return multierror.Append(domain.ErrServiceError, errors.New("unexpected postgres error"), err)
	}

	switch pgErr.Code {

	case pgerrcode.NotNullViolation:
		field := pgErr.ColumnName
		if field == "" {
			field = "field"
		}
		return multierror.Append(domain.ErrInvalid, errors.New(field+" is required"), err)

	case pgerrcode.CheckViolation:
		return multierror.Append(domain.ErrInvalid, errors.New("check constraint failed: "+pgErr.ConstraintName), err)

	case pgerrcode.InvalidTextRepresentation:
		return multierror.Append(domain.ErrInvalid, errors.New("invalid text representation"), err)

	case pgerrcode.StringDataRightTruncationDataException:
		field := pgErr.ColumnName
		if field == "" {
			field = "field"
		}
		return multierror.Append(domain.ErrInvalid, errors.New(field+" too long"), err)

	// --- CONFLICT ---
	case pgerrcode.UniqueViolation:
		return multierror.Append(domain.ErrExists, errors.New(pgErr.ConstraintName+" already exists"), err)

	case pgerrcode.ForeignKeyViolation:
		return multierror.Append(domain.ErrInvalid, errors.New("referenced row not found"), err)

	case pgerrcode.DeadlockDetected:
		return multierror.Append(domain.ErrConflict, errors.New("deadlock detected"), err)

	case pgerrcode.SerializationFailure:
		return multierror.Append(domain.ErrConflict, errors.New("serialization failed"), err)

	case
		pgerrcode.ConnectionFailure,
		pgerrcode.TooManyConnections, pgerrcode.AdminShutdown,
		pgerrcode.CrashShutdown,
		pgerrcode.CannotConnectNow:
		return multierror.Append(domain.ErrUnavailable, errors.New("database connection failed"), err)

	default:
		return multierror.Append(domain.ErrServiceError, errors.New("unhandled postgres error"), err)
	}
}
