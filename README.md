### Note : username and password in the code were created specifically for the hackathon and are now invalid or expired.

# Ayoba Business Chat API Hackathon Project Documentation
Project Overview
This document provides a detailed overview of the Golang server developed for the Ayoba Business Chat API Hackathon. The server was designed to interact with the Ayoba platform, sending and receiving messages to an Ayoba app user.

## Server Architecture
The server is built using the Go programming language and utilizes the Ayoba Business Chat API to communicate with the Ayoba platform. The primary components of the server include:

## HTTP Server: Handles incoming requests from the Ayoba platform and processes outgoing messages.
Ayoba API Client: Interacts with the Ayoba API to send and receive messages.
Message Handling Logic: Processes incoming messages and generates appropriate responses.
Server Functionality
The server performs the following key functions:

## Message Sending:

Receives a message from the user interface or other source.
Formats the message according to the Ayoba API requirements.
Sends the message to the specified Ayoba user using the Ayoba API.
Handles potential errors or failures during the sending process.

## Message Receiving:

Listens for incoming messages from the Ayoba platform.
Parses the received messages and extracts relevant information.
Processes the messages based on their content and type.
Responds to the sender if necessary.

## API Endpoints
The server exposes the following API endpoints:

### POST /sendtext: Accepts a JSON payload containing the message content and recipient's Ayoba ID.
### GET /gettext: Retrieves any unread messages from the Ayoba platform.


API Endpoints
The server exposes the following API endpoints:

POST /send_message: Accepts a JSON payload containing the message content and recipient's Ayoba ID.
GET /receive_messages: Retrieves any unread messages from the Ayoba platform.
