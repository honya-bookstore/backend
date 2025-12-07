INSERT INTO medium (id, url, alt_text) VALUES
('00000000-0000-8000-0000-000000000001', 'https://example.com/images/book1-cover.jpg', 'The Art of Programming cover'),
('00000000-0000-8000-0000-000000000002', 'https://example.com/images/book2-cover.jpg', 'Data Structures Explained cover'),
('00000000-0000-8000-0000-000000000003', 'https://example.com/images/book3-cover.jpg', 'Web Development Mastery cover'),
('00000000-0000-8000-0000-000000000004', 'https://example.com/images/book4-cover.jpg', 'Database Design Patterns cover'),
('00000000-0000-8000-0000-000000000005', 'https://example.com/images/book5-cover.jpg', 'Machine Learning Basics cover'),
('00000000-0000-8000-0000-000000000006', 'https://example.com/images/book6-cover.jpg', 'Cloud Architecture cover'),
('00000000-0000-8000-0000-000000000007', 'https://example.com/images/book7-cover.jpg', 'DevOps Handbook cover'),
('00000000-0000-8000-0000-000000000008', 'https://example.com/images/book8-cover.jpg', 'Security Fundamentals cover'),
('00000000-0000-8000-0000-000000000009', 'https://example.com/images/book9-cover.jpg', 'Mobile App Development cover'),
('00000000-0000-8000-0000-000000000010', 'https://example.com/images/book10-cover.jpg', 'API Design Principles cover'),
('00000000-0000-8000-0000-000000000011', 'https://example.com/images/book1-preview.jpg', 'The Art of Programming preview'),
('00000000-0000-8000-0000-000000000012', 'https://example.com/images/book3-preview.jpg', 'Web Development Mastery preview')
ON CONFLICT (id) DO NOTHING;

-- Seed Categories (5)
INSERT INTO categories (id, slug, name, description) VALUES
('00000000-0000-9000-0000-000000000001', 'programming', 'Programming', 'Books about programming languages and software development'),
('00000000-0000-9000-0000-000000000002', 'data-science', 'Data Science', 'Books about data analysis, machine learning, and AI'),
('00000000-0000-9000-0000-000000000003', 'web-development', 'Web Development', 'Books about frontend and backend web technologies'),
('00000000-0000-9000-0000-000000000004', 'devops', 'DevOps & Cloud', 'Books about deployment, CI/CD, and cloud infrastructure'),
('00000000-0000-9000-0000-000000000005', 'security', 'Security', 'Books about cybersecurity and secure coding practices')
ON CONFLICT (id) DO NOTHING;

-- Seed Books (10)
INSERT INTO books (id, title, description, author, price, pages_count, year, publisher, weight, stock_quantity, purchase_count, rating) VALUES
('00000000-0000-a000-0000-000000000001', 'The Art of Programming', 'A comprehensive guide to writing clean and efficient code.', 'John Smith', 450000, 520, 2023, 'Tech Books Publishing', 0.8, 50, 120, 4.5),
('00000000-0000-a000-0000-000000000002', 'Data Structures Explained', 'Master data structures with practical examples.', 'Jane Doe', 380000, 420, 2022, 'Code Academy Press', 0.7, 35, 89, 4.2),
('00000000-0000-a000-0000-000000000003', 'Web Development Mastery', 'From HTML to modern frameworks - a complete journey.', 'Michael Chen', 520000, 680, 2024, 'WebDev Publications', 1.1, 75, 210, 4.8),
('00000000-0000-a000-0000-000000000004', 'Database Design Patterns', 'Best practices for designing scalable databases.', 'Sarah Johnson', 420000, 380, 2023, 'Data Systems Inc', 0.6, 40, 65, 4.0),
('00000000-0000-a000-0000-000000000005', 'Machine Learning Basics', 'An introduction to ML algorithms and applications.', 'David Wilson', 580000, 550, 2024, 'AI Research Press', 0.9, 60, 180, 4.6),
('00000000-0000-a000-0000-000000000006', 'Cloud Architecture', 'Design patterns for cloud-native applications.', 'Emily Brown', 490000, 450, 2023, 'Cloud Computing Books', 0.75, 45, 95, 4.3),
('00000000-0000-a000-0000-000000000007', 'DevOps Handbook', 'Implementing CI/CD and automation best practices.', 'Robert Taylor', 410000, 400, 2022, 'Agile Publishers', 0.65, 55, 140, 4.4),
('00000000-0000-a000-0000-000000000008', 'Security Fundamentals', 'Essential cybersecurity concepts for developers.', 'Lisa Anderson', 360000, 320, 2024, 'SecureCode Press', 0.5, 30, 75, 4.1),
('00000000-0000-a000-0000-000000000009', 'Mobile App Development', 'Build cross-platform apps with modern tools.', 'Kevin Martinez', 470000, 510, 2023, 'Mobile Tech Books', 0.85, 65, 155, 4.5),
('00000000-0000-a000-0000-000000000010', 'API Design Principles', 'Create robust and scalable APIs.', 'Amanda Lee', 340000, 290, 2024, 'Backend Publishing', 0.45, 40, 88, 4.2)
ON CONFLICT (id) DO NOTHING;

-- Seed Books Medium (linking books to their images)
INSERT INTO books_medium (book_id, media_id, "order", is_cover) VALUES
('00000000-0000-a000-0000-000000000001', '00000000-0000-8000-0000-000000000001', 0, TRUE),
('00000000-0000-a000-0000-000000000001', '00000000-0000-8000-0000-000000000011', 1, FALSE),
('00000000-0000-a000-0000-000000000002', '00000000-0000-8000-0000-000000000002', 0, TRUE),
('00000000-0000-a000-0000-000000000003', '00000000-0000-8000-0000-000000000003', 0, TRUE),
('00000000-0000-a000-0000-000000000003', '00000000-0000-8000-0000-000000000012', 1, FALSE),
('00000000-0000-a000-0000-000000000004', '00000000-0000-8000-0000-000000000004', 0, TRUE),
('00000000-0000-a000-0000-000000000005', '00000000-0000-8000-0000-000000000005', 0, TRUE),
('00000000-0000-a000-0000-000000000006', '00000000-0000-8000-0000-000000000006', 0, TRUE),
('00000000-0000-a000-0000-000000000007', '00000000-0000-8000-0000-000000000007', 0, TRUE),
('00000000-0000-a000-0000-000000000008', '00000000-0000-8000-0000-000000000008', 0, TRUE),
('00000000-0000-a000-0000-000000000009', '00000000-0000-8000-0000-000000000009', 0, TRUE),
('00000000-0000-a000-0000-000000000010', '00000000-0000-8000-0000-000000000010', 0, TRUE)
ON CONFLICT (book_id, media_id) DO NOTHING;

-- Seed Books Categories (linking books to categories)
INSERT INTO books_categories (book_id, category_id) VALUES
-- The Art of Programming -> Programming
('00000000-0000-a000-0000-000000000001', '00000000-0000-9000-0000-000000000001'),
-- Data Structures Explained -> Programming, Data Science
('00000000-0000-a000-0000-000000000002', '00000000-0000-9000-0000-000000000001'),
('00000000-0000-a000-0000-000000000002', '00000000-0000-9000-0000-000000000002'),
-- Web Development Mastery -> Web Development, Programming
('00000000-0000-a000-0000-000000000003', '00000000-0000-9000-0000-000000000003'),
('00000000-0000-a000-0000-000000000003', '00000000-0000-9000-0000-000000000001'),
-- Database Design Patterns -> Programming, Data Science
('00000000-0000-a000-0000-000000000004', '00000000-0000-9000-0000-000000000001'),
('00000000-0000-a000-0000-000000000004', '00000000-0000-9000-0000-000000000002'),
-- Machine Learning Basics -> Data Science
('00000000-0000-a000-0000-000000000005', '00000000-0000-9000-0000-000000000002'),
-- Cloud Architecture -> DevOps & Cloud
('00000000-0000-a000-0000-000000000006', '00000000-0000-9000-0000-000000000004'),
-- DevOps Handbook -> DevOps & Cloud
('00000000-0000-a000-0000-000000000007', '00000000-0000-9000-0000-000000000004'),
-- Security Fundamentals -> Security, Programming
('00000000-0000-a000-0000-000000000008', '00000000-0000-9000-0000-000000000005'),
('00000000-0000-a000-0000-000000000008', '00000000-0000-9000-0000-000000000001'),
-- Mobile App Development -> Programming, Web Development
('00000000-0000-a000-0000-000000000009', '00000000-0000-9000-0000-000000000001'),
('00000000-0000-a000-0000-000000000009', '00000000-0000-9000-0000-000000000003'),
-- API Design Principles -> Web Development, DevOps
('00000000-0000-a000-0000-000000000010', '00000000-0000-9000-0000-000000000003'),
('00000000-0000-a000-0000-000000000010', '00000000-0000-9000-0000-000000000004')
ON CONFLICT (book_id, category_id) DO NOTHING;

-- Seed Reviews
INSERT INTO reviews (id, rating, vote_count, content, user_id, book_id) VALUES
('00000000-0000-b000-0000-000000000001', 5, 12, 'Excellent book! Really helped me understand clean code principles.', '00000000-0000-7000-0000-000000000001', '00000000-0000-a000-0000-000000000001'),
('00000000-0000-b000-0000-000000000002', 4, 8, 'Great content but could use more advanced examples.', '00000000-0000-7000-0000-000000000001', '00000000-0000-a000-0000-000000000002'),
('00000000-0000-b000-0000-000000000003', 5, 25, 'The best web development book I have ever read!', '00000000-0000-7000-0000-000000000001', '00000000-0000-a000-0000-000000000003'),
('00000000-0000-b000-0000-000000000004', 4, 6, 'Very practical approach to database design.', '00000000-0000-7000-0000-000000000001', '00000000-0000-a000-0000-000000000004'),
('00000000-0000-b000-0000-000000000005', 5, 18, 'Perfect introduction to machine learning concepts.', '00000000-0000-7000-0000-000000000001', '00000000-0000-a000-0000-000000000005'),
('00000000-0000-b000-0000-000000000006', 4, 10, 'Solid coverage of cloud architecture patterns.', '00000000-0000-7000-0000-000000000001', '00000000-0000-a000-0000-000000000006'),
('00000000-0000-b000-0000-000000000007', 5, 15, 'A must-read for anyone getting into DevOps.', '00000000-0000-7000-0000-000000000001', '00000000-0000-a000-0000-000000000007'),
('00000000-0000-b000-0000-000000000008', 3, 4, 'Good basics but lacks depth in some areas.', '00000000-0000-7000-0000-000000000001', '00000000-0000-a000-0000-000000000008')
ON CONFLICT (id) DO NOTHING;

-- Seed Review Votes
INSERT INTO review_votes (id, user_id, review_id, is_up) VALUES
('00000000-0000-c000-0000-000000000001', '00000000-0000-7000-0000-000000000001', '00000000-0000-b000-0000-000000000001', TRUE),
('00000000-0000-c000-0000-000000000002', '00000000-0000-7000-0000-000000000001', '00000000-0000-b000-0000-000000000002', TRUE),
('00000000-0000-c000-0000-000000000003', '00000000-0000-7000-0000-000000000001', '00000000-0000-b000-0000-000000000003', TRUE),
('00000000-0000-c000-0000-000000000004', '00000000-0000-7000-0000-000000000001', '00000000-0000-b000-0000-000000000005', TRUE),
('00000000-0000-c000-0000-000000000005', '00000000-0000-7000-0000-000000000001', '00000000-0000-b000-0000-000000000007', TRUE),
('00000000-0000-c000-0000-000000000006', '00000000-0000-7000-0000-000000000001', '00000000-0000-b000-0000-000000000008', FALSE)
ON CONFLICT (id) DO NOTHING;

-- Carts
INSERT INTO carts (id, user_id) VALUES
('00000000-0000-7000-0000-000000000001', '00000000-0000-7000-0000-000000000006')
 ON CONFLICT (id) DO NOTHING;

-- Cart Items
INSERT INTO cart_items (id, cart_id, book_id, quantity) VALUES
('00000000-0000-7000-0000-000000000001', '00000000-0000-7000-0000-000000000001', '00000000-0000-a000-0000-000000000001', 10),
('00000000-0000-7000-0000-000000000002', '00000000-0000-7000-0000-000000000001', '00000000-0000-a000-0000-000000000002', 1)
 ON CONFLICT (id) DO NOTHING;
