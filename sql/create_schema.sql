CREATE table questions (
    `id` INT(10) NOT NULL AUTO_INCREMENT,
    `text` VARCHAR(1000) NULL DEFAULT NULL,
    `created` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

CREATE table answers (
    `id` INT(10) NOT NULL AUTO_INCREMENT,
    `text` VARCHAR(1000) NOT NULL,
    `correct` TINYINT(1) NULL DEFAULT 0,
    `question_id` INT(10) NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (question_id) REFERENCES questions(id) ON DELETE CASCADE
);