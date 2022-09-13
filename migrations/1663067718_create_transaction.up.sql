CREATE TABLE `Transactions`(
    `id` VARCHAR(40) NOT NULL,
    `issuedate` VARCHAR(40) NOT NULL,
    `duedate` VARCHAR(40) NOT NULL,
    `actualreturndate` VARCHAR(40) NOT NULL,
    `book_id` VARCHAR(40) NOT NULL,
    `user_id` VARCHAR(40) NOT NULL,
    PRIMARY KEY(id),
    FOREIGN KEY(user_id) REFERENCES user(id),
    FOREIGN KEY(book_id) REFERENCES book(id)
);