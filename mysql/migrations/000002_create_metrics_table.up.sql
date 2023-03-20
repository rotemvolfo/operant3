CREATE TABLE IF NOT EXISTS metrics(
    id INT NOT NULL UNIQUE AUTO_INCREMENT,
    name VARCHAR (255) NOT NULL UNIQUE,
    high_treshold INT ,
    low_treshold INT ,
    current INT ,
    user_id VARCHAR (255) ,
    FOREIGN KEY (user_id) REFERENCES services(user_id) ,
    PRIMARY KEY (id)
)