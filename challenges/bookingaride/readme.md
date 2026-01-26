# Coding Kata Challenge — Booking Service (Backend)

 You’re building a **Booking Service** for a premium chauffeured rides platform.
The service is responsible for creating and managing bookings and supporting a small set of booking operations used by:
- a guest-facing app (customers),
- an internal operations tool,
- and downstream services (dispatching / payments) that will integrate later.

This kata is **not** about algorithms. It’s about:
- clean backend design,
- correct business rules,
- test coverage,
- and how you work under constraints.

Timebox: **90 minutes**.


## My solutions

guess chauffeur ride booking datetime location
request estimation
create a booking
get booking
cancel booking

who can create -> user - token/jwt

### Guess Ops
POST /booking
Headers
X-Idemportency-Key: UUID
Authentication: Bearer XXX
Body
{
    "guess_id": UUID,
    "pickup_at": "2026-1-21T000:00:00Z",
    "pickup": {"lat": 12.35, "lng": 12.45},
    "dropoff": {"lat": 12.35, "lng": 12.45}
}
Response
{
    "booking_id": "UUID"
    "guess_id": UUID,
    "pickup_at": "2026-1-21T000:00:00Z",
    "pickup": {"lat": 12.35, "lng": 12.45},
    "dropoff": {"lat": 12.35, "lng": 12.45}
    "created_at": "2026-1-21T000:00:00Z"
    "status": "PENDING"
}
Create a Booking with status: PENDING.
RequestID must be unique 

### Rider Ops

POST /booking/{id}/confirm
Authentication: Bearer XXX

Reponse
OK 200 
Conflict 409

Change the status from PENDING to CONFIRMED
Lock row with SELECT * FOR UPDATE
If status == PENDING -> update status and 
If status != PENDING -> conflict






### Entity

Booking
ID
coord: lat,long
status: PENDING, CONFIRMED, CANCELED, REJECTED
UserID
