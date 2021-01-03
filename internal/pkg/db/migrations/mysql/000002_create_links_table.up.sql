CREATE TABLE IF NOT EXISTS Links(
    ID uuid DEFAULT uuid_generate_v4(),
    Title VARCHAR (255) ,
    Address VARCHAR (255) ,
    UserID uuid ,
    FOREIGN KEY (UserID) REFERENCES Users(ID) ,
    PRIMARY KEY (ID)
)