On to the exercise - we will start by creating a database along with a phone_numbers table. Inside that table we want to add the following entries (yes, I know there are duplicates):

1234567890  
123 456 7891  
(123) 456 7892  
(123) 456-7893  
123-456-7894  
123-456-7890  
1234567892  
(123)456-7892  

Once you have all the data in the DB, our next step is to normalize the phone number. We are going to update all of our numbers so that they match the format:

##########

Once you written code that will successfully take a number in with any format and return the same number in the proper format we are going to use an UPDATE to alter the entries in the database. If the value we are inserting into our database already exists (it is a duplicate), we will instead be deleting the original entry.

When your program is done your database entries should look like this (the order is irrelevant, but duplicates should be removed):

1234567890
1234567891
1234567892
1234567893
1234567894
