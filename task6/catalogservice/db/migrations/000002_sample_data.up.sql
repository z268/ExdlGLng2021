START TRANSACTION;

INSERT INTO `authors` (`uuid`, `name`, `created_at`, `updated_at`, `deleted_at`) VALUES
('1dfe4bd9-3a71-4761-bc68-6e9c040ba67f', 'George Ornbo', '2021-06-02 06:04:22', '2021-06-02 06:04:22', NULL),
('4f4716af-0672-491e-8898-82d5e783070c', 'Jay McGavren', '2021-06-02 06:04:10', '2021-06-02 06:04:10', NULL),
('62c5bf57-b998-4f8b-ac9d-48df49fca0f6', 'Caleb Doxsey', '2021-06-02 06:04:30', '2021-06-02 06:04:30', NULL),
('680bb29b-2271-4feb-a625-41da44bb6dc5', 'Alan A. A. Donovan', '2021-06-02 06:04:00', '2021-06-02 06:04:00', NULL),
('895dd59e-9978-4381-9888-7adaacc3543b', 'Al Sweigart', '2021-06-02 05:59:00', '2021-06-02 05:59:00', NULL),
('97fed32f-8f04-42c9-88ed-a16706c13c00', 'Mark Lutz', '2021-06-02 06:00:12', '2021-06-02 06:00:12', NULL),
('aa74c326-158f-4ceb-a73b-6e1d9fc16a08', 'Anthony Scopatz', '2021-06-02 05:59:21', '2021-06-02 05:59:21', NULL),
('be558f6f-a9b2-49e3-8a35-9f97c90db55d', 'Allen B. Downey', '2021-06-02 05:59:11', '2021-06-02 05:59:11', NULL),
('c3401366-3334-4119-a772-d9f1f48e16b0', 'Nathan Youngman', '2021-06-02 06:04:37', '2021-06-02 06:04:37', NULL),
('e404691e-626d-493b-9688-bb801c93b78b', 'Eric Matthes', '2021-06-02 05:58:43', '2021-06-02 05:58:43', NULL),
('e6a75705-407c-4170-b52b-0278b1556493', 'Joanne Roling', '2021-05-23 19:06:56', '2021-05-23 19:06:56', NULL),
('ed683819-4039-493e-97e3-75b6efc674dd', 'Paul Barry', '2021-06-02 05:58:51', '2021-06-02 05:58:51', NULL);

INSERT INTO `books` (`uuid`, `name`, `created_at`, `updated_at`, `deleted_at`, `author_uuid`) VALUES
('1c2626d7-40ba-4697-9c7f-4a99122378e3', 'Automate the Boring Stuff with Python, 2nd Edition: Practical Programming for Total Beginners', '2021-06-02 06:07:30', '2021-06-02 06:07:30', NULL, '895dd59e-9978-4381-9888-7adaacc3543b'),
('375605c6-8156-4754-8736-e4acb1ffdad0', 'Python Crash Course', '2021-06-02 06:06:55', '2021-06-02 06:06:55', NULL, 'e404691e-626d-493b-9688-bb801c93b78b'),
('65a6bf6f-1673-420d-9ecd-270b06a5e48c', 'Head-First Python (2nd edition)', '2021-06-02 06:08:03', '2021-06-02 06:08:03', NULL, 'ed683819-4039-493e-97e3-75b6efc674dd'),
('6a81d70d-19ac-4b09-8794-d75c3b5822b9', 'Sams Teach Yourself Go in 24 Hours: Next Generation Systems Programming with Golang', '2021-06-02 06:10:16', '2021-06-02 06:10:16', NULL, '1dfe4bd9-3a71-4761-bc68-6e9c040ba67f'),
('7b6d5e3a-ab6c-43cc-8794-68d62eb85a11', 'Learning with Python: How to Think Like a Computer Scientist', '2021-06-02 06:08:42', '2021-06-02 06:08:42', NULL, 'be558f6f-a9b2-49e3-8a35-9f97c90db55d'),
('a7951781-d0c4-4eab-b32a-833fea587a11', 'Learning Python, 5th Edition', '2021-06-02 06:34:06', '2021-06-02 06:34:06', NULL, '97fed32f-8f04-42c9-88ed-a16706c13c00'),
('dd1280ea-284e-4944-939e-ef36c3cd83ad', 'Get Programming with Go', '2021-06-02 06:10:53', '2021-06-02 06:10:53', NULL, 'c3401366-3334-4119-a772-d9f1f48e16b0'),
('dee5a189-0843-42be-8d75-8cca4e536e34', 'Introducing Go: Build Reliable, Scalable Programs', '2021-06-02 06:10:36', '2021-06-02 06:10:36', NULL, '62c5bf57-b998-4f8b-ac9d-48df49fca0f6'),
('eae533ff-12c6-4240-8da7-9b0775930a63', 'Head First Go', '2021-06-02 06:09:44', '2021-06-02 06:09:44', NULL, '4f4716af-0672-491e-8898-82d5e783070c'),
('f0343716-4d43-46e7-b7fa-049e03bc7f91', 'The Go Programming Language', '2021-06-02 06:09:24', '2021-06-02 06:09:24', NULL, '680bb29b-2271-4feb-a625-41da44bb6dc5');

INSERT INTO `categories` (`uuid`, `name`, `created_at`, `updated_at`, `deleted_at`, `parent_uuid`) VALUES
('524ab1ae-3293-4e90-9458-88b61eb95f0d', 'Software development', '2021-05-23 19:07:25', '2021-06-02 05:53:48', NULL, NULL),
('7d2b4f0b-daa7-4471-9067-11a1ab46beca', 'Python', '2021-05-23 19:07:25', '2021-06-02 05:53:53', NULL, '524ab1ae-3293-4e90-9458-88b61eb95f0d'),
('b35d3107-9df1-4058-86fc-c875337377d0', 'Golang', '2021-06-02 05:55:47', '2021-06-02 05:58:01', NULL, '524ab1ae-3293-4e90-9458-88b61eb95f0d');

INSERT INTO `book_categories` (`book_uuid`, `category_uuid`, `created_at`) VALUES
('1c2626d7-40ba-4697-9c7f-4a99122378e3', '7d2b4f0b-daa7-4471-9067-11a1ab46beca', '2021-06-02 06:14:07'),
('375605c6-8156-4754-8736-e4acb1ffdad0', '7d2b4f0b-daa7-4471-9067-11a1ab46beca', '2021-06-02 06:14:07'),
('65a6bf6f-1673-420d-9ecd-270b06a5e48c', '7d2b4f0b-daa7-4471-9067-11a1ab46beca', '2021-06-02 06:14:07'),
('6a81d70d-19ac-4b09-8794-d75c3b5822b9', 'b35d3107-9df1-4058-86fc-c875337377d0', '2021-06-02 06:15:54'),
('7b6d5e3a-ab6c-43cc-8794-68d62eb85a11', '7d2b4f0b-daa7-4471-9067-11a1ab46beca', '2021-06-02 06:14:07'),
('a7951781-d0c4-4eab-b32a-833fea587a11', '7d2b4f0b-daa7-4471-9067-11a1ab46beca', '2021-06-02 06:39:03'),
('dd1280ea-284e-4944-939e-ef36c3cd83ad', 'b35d3107-9df1-4058-86fc-c875337377d0', '2021-06-02 06:15:54'),
('dee5a189-0843-42be-8d75-8cca4e536e34', 'b35d3107-9df1-4058-86fc-c875337377d0', '2021-06-02 06:15:54'),
('eae533ff-12c6-4240-8da7-9b0775930a63', 'b35d3107-9df1-4058-86fc-c875337377d0', '2021-06-02 06:15:54'),
('f0343716-4d43-46e7-b7fa-049e03bc7f91', 'b35d3107-9df1-4058-86fc-c875337377d0', '2021-06-02 06:15:54');

COMMIT;