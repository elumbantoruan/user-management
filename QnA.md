How long did this assignment take?
> Around 5 hours

What was the hardest part?
> That the was time for me to use GORM in Golang.  Most of work these days involve with NoSQL.  But it was a pleasant to know the power of postgres and GORM.

Did you learn anything new?
> Yes, the GORM in postgres

Is there anything you would have liked to implement but didn't have the time to?
> I would like to use PGCRYPTO, but it will definitely take more time from my end
> I would like to include database script to create the database and to use the extension PGCRYPTO

What are the security holes (if any) in your system? If there are any, how would you fix them?
> I store the password as SHA256 plus Salt.  People always mentioned do not roll your own crypto.  I should be using PGCRYPTO

Do you feel that your skills were well tested?
> I do.  I like this, it does require attention to detail such as unique field (email) in database, adding claims in JWT to store email, validating JWT token
