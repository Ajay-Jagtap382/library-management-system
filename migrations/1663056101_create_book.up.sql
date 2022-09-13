CREATE TABLE book
(
    `id` VARCHAR(64) NOT NULL,
    `bookName` VARCHAR(20) NOT NULL,
    `description` VARCHAR(200) NOT NULL,
    `totalCopies` INT NOT NULL,
    `currentCopies` INT NOT NULL,
    PRIMARY KEY (`id`)
);