START TRANSACTION;

CREATE TABLE `authors` (
                           `uuid` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                           `name` varchar(255) NOT NULL,
                           `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                           `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                           `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `books` (
                         `uuid` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                         `name` varchar(255) NOT NULL,
                         `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         `deleted_at` datetime DEFAULT NULL,
                         `author_uuid` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `book_categories` (
                                   `book_uuid` varchar(40) NOT NULL,
                                   `category_uuid` varchar(40) NOT NULL,
                                   `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `categories` (
                              `uuid` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                              `name` varchar(255) NOT NULL,
                              `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                              `deleted_at` datetime DEFAULT NULL,
                              `parent_uuid` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


ALTER TABLE `authors`
    ADD PRIMARY KEY (`uuid`);

ALTER TABLE `books`
    ADD PRIMARY KEY (`uuid`),
    ADD KEY `books_to_author` (`author_uuid`);

ALTER TABLE `book_categories`
    ADD UNIQUE KEY `book_category` (`book_uuid`,`category_uuid`) USING BTREE,
    ADD KEY `category_uuid` (`category_uuid`);

ALTER TABLE `categories`
    ADD PRIMARY KEY (`uuid`),
    ADD KEY `parent_uuid` (`parent_uuid`);


ALTER TABLE `books`
    ADD CONSTRAINT `books_to_author` FOREIGN KEY (`author_uuid`) REFERENCES `authors` (`uuid`) ON DELETE RESTRICT ON UPDATE RESTRICT;

ALTER TABLE `book_categories`
    ADD CONSTRAINT `author` FOREIGN KEY (`category_uuid`) REFERENCES `categories` (`uuid`) ON DELETE CASCADE ON UPDATE RESTRICT,
    ADD CONSTRAINT `book` FOREIGN KEY (`book_uuid`) REFERENCES `books` (`uuid`) ON DELETE CASCADE ON UPDATE RESTRICT;

ALTER TABLE `categories`
    ADD CONSTRAINT `parent_uuid` FOREIGN KEY (`parent_uuid`) REFERENCES `categories` (`uuid`) ON DELETE RESTRICT;

COMMIT;