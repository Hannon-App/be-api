
# Hannon.Apps

![Logo](https://github.com/Hannon-App/be-api/blob/main/Hannon.png)

Readme in english [click here](https://github.com/Hannon-App/be-api/blob/main/README_english.md).

Hannon apps is an application for borrowing camping equipment/outdoor equipment attributes that bridges users (potential borrowers) with tenants (loan parties). With this application, it will be easier and more flexible for tenants to run their outdoor equipment rental business, and from the user side (potential borrowers) it will be easier to find items available from tenants at affordable prices because there is no need to buy/have equipment when carrying out activities. outdoors (camping, hiking, etc.). Apart from that, with the Hannon App in the future, it can increase the intensity of tourism so that it will have a good impact on tourism development, thus revitalizing economic activity in the region.


## Fitur Users 

- Registers
- Login Users
- Edit Users
- view items & tenants
- Users search for items from tenants according to their destination city
- Users can borrow items from tenants according to the time specified by the user.
- Users can cancel loans that have been booked with valid reasons
- Users can make payments via Xendit which has been integrated
- Users can receive payment notifications via active email.

## Fitur Tenants

- Register
- Tenants Login
- Tenants can insert their products to be loaned
- Tenants can edit/update their items product data
- Tenants can delete their item product data
- Tenants can make transactions with users

## Open APIs

For Open API, you can see more details [click here](https://github.com/Hannon-App/be-api/blob/main/hannonapp-openAPI.yml)

## ERD

Untuk ERD Diagram [click here](https://app.diagrams.net/#G1vUnt4shuvShWc86VFJCGuL399IyZq0Bi)


## Running Local
Cloning project

```bash
  $ https://github.com/Hannon-App/be-api.git
```

Go to the project directory

```bash
  $ cd ~/your project name
```
Create a new `database`

Create a file with the name in the project root folder `.env` with the format below. Adjust the configuration on the local computer

```bash
export DBUSER='root'
export DBPASS='enter your password'
export DBHOST='127.0.0.1'
export DBPORT='3306'
export DBNAME='your database name'
export JWTSECRET='......'
export KEY_API='......'
export KEY_API_SECRET='.........'
export CLOUD_NAME='.....'
export GOOGLE_APPLICATION_CREDENTIALS='keys.json'
export XENDIT_SECRET_KEY='enter secret key from xendit'
export CALLBACK_KEY='enter the callback key from xendit'
```

Run the application

```bash
  $ go run main.go
```


## Authors

- [@firhanaf](https://github.com/firhanaf)
- [@royanqodri](https://github.com/royanqodri)
- [@Prayogarock](https://github.com/Prayogarock)

 
