CREATE TABLE IF NOT EXISTS todos (
    id            BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    content          VARCHAR(255) NOT NULL,
    checked       BOOLEAN NOT NULL,
    removed       BOOLEAN NOT NULL
)
