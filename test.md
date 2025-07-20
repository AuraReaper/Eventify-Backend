# API Testing Guide

Base URL: `http://localhost:3000`

## Authentication Endpoints

### Register User
- **URL**: `/api/auth/register`
- **Method**: `POST`
- **Auth Required**: No
- **Body**:
```json
{
    "email": "user@example.com",
    "password": "yourpassword",
    "role": "attendee"  // or "manager"
}
```
- **Success Response** (201):
```json
{
    "status": "success",
    "message": "User registered successfully",
    "data": {
        "user": {
            "id": 1,
            "email": "user@example.com",
            "role": "attendee"
        },
        "token": "jwt_token_here"
    }
}
```

### Login
- **URL**: `/api/auth/login`
- **Method**: `POST`
- **Auth Required**: No
- **Body**:
```json
{
    "email": "user@example.com",
    "password": "yourpassword"
}
```
- **Success Response** (200):
```json
{
    "status": "success",
    "message": "Login successful",
    "data": {
        "user": {
            "id": 1,
            "email": "user@example.com",
            "role": "attendee"
        },
        "token": "jwt_token_here"
    }
}
```

## Event Endpoints

### Get All Events
- **URL**: `/api/events`
- **Method**: `GET`
- **Auth Required**: Yes
- **Headers**: `Authorization: Bearer <token>`
- **Success Response** (200):
```json
{
    "status": "success",
    "message": "Get all event success",
    "data": [
        {
            "id": 1,
            "title": "Event Title",
            "description": "Event Description",
            "date": "2024-03-20T00:00:00Z",
            "location": "Event Location",
            "user_id": 1
        }
    ]
}
```

### Get Event by ID
- **URL**: `/api/events/:eventId`
- **Method**: `GET`
- **Auth Required**: Yes
- **Headers**: `Authorization: Bearer <token>`
- **Success Response** (200):
```json
{
    "status": "success",
    "message": "Get event by id success",
    "data": {
        "id": 1,
        "title": "Event Title",
        "description": "Event Description",
        "date": "2024-03-20T00:00:00Z",
        "location": "Event Location",
        "user_id": 1
    }
}
```

### Create Event
- **URL**: `/api/events`
- **Method**: `POST`
- **Auth Required**: Yes
- **Headers**: `Authorization: Bearer <token>`
- **Body**:
```json
{
    "title": "Event Title",
    "description": "Event Description",
    "date": "2024-03-20T00:00:00Z",
    "location": "Event Location"
}
```
- **Success Response** (201):
```json
{
    "status": "success",
    "message": "Create event success",
    "data": {
        "id": 1,
        "title": "Event Title",
        "description": "Event Description",
        "date": "2024-03-20T00:00:00Z",
        "location": "Event Location",
        "user_id": 1
    }
}
```

### Update Event
- **URL**: `/api/events/:eventId`
- **Method**: `PUT`
- **Auth Required**: Yes
- **Headers**: `Authorization: Bearer <token>`
- **Body**:
```json
{
    "title": "Updated Title",
    "description": "Updated Description"
}
```
- **Success Response** (200):
```json
{
    "status": "success",
    "message": "Update event success",
    "data": {
        "id": 1,
        "title": "Updated Title",
        "description": "Updated Description",
        "date": "2024-03-20T00:00:00Z",
        "location": "Event Location",
        "user_id": 1
    }
}
```

### Delete Event
- **URL**: `/api/events/:eventId`
- **Method**: `DELETE`
- **Auth Required**: Yes
- **Headers**: `Authorization: Bearer <token>`
- **Success Response** (200):
```json
{
    "status": "success",
    "message": "Delete event success"
}
```

## Ticket Endpoints

### Get All Tickets
- **URL**: `/api/tickets`
- **Method**: `GET`
- **Auth Required**: Yes
- **Headers**: `Authorization: Bearer <token>`
- **Success Response** (200):
```json
{
    "status": "success",
    "message": "Get all ticket success",
    "data": [
        {
            "id": 1,
            "event_id": 1,
            "user_id": 1,
            "entered": false,
            "event": {
                "id": 1,
                "title": "Event Title",
                "description": "Event Description",
                "date": "2024-03-20T00:00:00Z",
                "location": "Event Location",
                "user_id": 1
            }
        }
    ]
}
```

### Get Ticket by ID
- **URL**: `/api/tickets/:ticketId`
- **Method**: `GET`
- **Auth Required**: Yes
- **Headers**: `Authorization: Bearer <token>`
- **Success Response** (200):
```json
{
    "status": "success",
    "message": "Get ticket by id success",
    "data": {
        "ticket": {
            "id": 1,
            "event_id": 1,
            "user_id": 1,
            "entered": false,
            "event": {
                "id": 1,
                "title": "Event Title",
                "description": "Event Description",
                "date": "2024-03-20T00:00:00Z",
                "location": "Event Location",
                "user_id": 1
            }
        },
        "QRCode": "base64_encoded_qr_code"
    }
}
```

### Create Ticket
- **URL**: `/api/tickets`
- **Method**: `POST`
- **Auth Required**: Yes (Manager only)
- **Headers**: `Authorization: Bearer <token>`
- **Body**:
```json
{
    "event_id": 1,
    "user_id": 2
}
```
- **Success Response** (201):
```json
{
    "status": "success",
    "message": "Create ticket success",
    "data": {
        "id": 1,
        "event_id": 1,
        "user_id": 2,
        "entered": false
    }
}
```

### Update Ticket
- **URL**: `/api/tickets/:ticketId`
- **Method**: `PUT`
- **Auth Required**: Yes
- **Headers**: `Authorization: Bearer <token>`
- **Body**:
```json
{
    "entered": true
}
```
- **Success Response** (200):
```json
{
    "status": "success",
    "message": "Update ticket success",
    "data": {
        "id": 1,
        "event_id": 1,
        "user_id": 1,
        "entered": true
    }
}
```

### Validate Ticket
- **URL**: `/api/tickets/validate`
- **Method**: `POST`
- **Auth Required**: Yes (Manager only)
- **Headers**: `Authorization: Bearer <token>`
- **Body**:
```json
{
    "ticket_id": 1
}
```
- **Success Response** (200):
```json
{
    "status": "success",
    "message": "Welcome to the event, please enter the event with your ticket id",
    "data": {
        "ticket": {
            "id": 1,
            "event_id": 1,
            "user_id": 1,
            "entered": true
        },
        "QRCode": "base64_encoded_qr_code"
    }
}
```

### Get Event Ticket Statistics
- **URL**: `/api/tickets/stats/:eventId`
- **Method**: `GET`
- **Auth Required**: Yes (Manager only)
- **Headers**: `Authorization: Bearer <token>`
- **Success Response** (200):
```json
{
    "status": "success",
    "message": "Get event ticket statistics success",
    "data": {
        "total_tickets": 100,
        "attended_tickets": 75,
        "attendance_rate": 75.0
    }
}
```

## Testing Notes

1. **Authentication**:
   - All endpoints except register and login require a valid JWT token
   - Include the token in the Authorization header: `Bearer <token>`

2. **Role-based Access**:
   - Only managers can create tickets
   - Only managers can validate tickets
   - Only managers can view ticket statistics
   - Regular users can only view and update their own tickets

3. **QR Code**:
   - QR codes are generated for ticket validation
   - QR codes contain ticket ID and owner ID information

4. **Error Responses**:
   - All endpoints return error responses in the format:
   ```json
   {
       "status": "fail",
       "message": "Error message here"
   }
   ```

5. **Testing Flow**:
   1. Register a manager account
   2. Register an attendee account
   3. Login as manager to create an event
   4. Create tickets for attendees
   5. Login as attendee to view their tickets
   6. Login as manager to validate tickets
   7. Check ticket statistics as manager 