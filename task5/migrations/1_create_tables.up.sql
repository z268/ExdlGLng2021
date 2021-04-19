CREATE TABLE `tasks` (
    `id` int NOT NULL,
    `name` varchar(255) DEFAULT NULL,
    `description` varchar(255) DEFAULT NULL,
    `duedate` datetime(6) DEFAULT NULL,
    `status` int DEFAULT '0'
);

ALTER TABLE `tasks`
    ADD PRIMARY KEY (`id`);

ALTER TABLE `tasks`
    MODIFY `id` int NOT NULL AUTO_INCREMENT;

COMMIT;