# Building Controllers

## Controller V1

Legacy way, minimal abstraction

```go
var uncheckedFields = map[string]bool{"UserID": true, "ID": true, "Description": true}

func createCategory(w http.ResponseWriter, r *http.Request) {
  var c models.Category
  err := json.NewDecoder(r.Body).Decode(&c)
  if BFE.HandleError(w, err) {
    return
  }

  err = models.IsModelValid(c, uncheckedFields)
  if BFE.HandleError(w, err) {
    return
  }

  claims, err := models.GetUserClaims(r)
  if BFE.HandleError(w, err) {
    return
  }

  c.UserId, _ = uuid.Parse(claims.UserID)
  err = DB.CreateCategory(&c)
  if BFE.HandleError(w, err) {
    return
  }

  RSP.SendObjectResponse(w, http.StatusCreated, c)
}
```

## Controller V2

Using models.GetUserID instead of getting the claims and then parsing the ID

Preferred usage of endpoints that don't have a request body

```go
var uncheckedFields = map[string]bool{"UserID": true, "SessionID": true, "Description": true, "Duration": true}

func createSession(w http.ResponseWriter, r *http.Request) {
  var t models.Toast
  err := json.NewDecoder(r.Body).Decode(&t)
  if BFE.HandleError(w, err) {
    return
  }

  err = models.IsModelValid(t, uncheckedFields)
  if BFE.HandleError(w, err) {
    return
  }

  t.UserID, err = models.GetUserID(r)
  if BFE.HandleError(w, err) {
    return
  }

  err = DB.CreateToastSession(&t)
  if BFE.HandleError(w, err) {
    return
  }

  RSP.SendObjectResponse(w, http.StatusCreated, t)
}
```

## Controller V3

Using FillModelFromJSON instead of parsing the model then validating it. This way you also can get the request fields for patch requests, and configure ForbiddenFields

Can't be used for controllers that don't fill up a model from a request like GET endpoints

```go
var config = models.ValidationConfig{
  IgnoreFields: map[string]bool{
    "Description": true,  // Optional field
  },
  ForbiddenFields: map[string]bool{
    "user_id": true,     // Set by server
    "session_id": true,  // Set by server
    "duration": true,    // Calculated on stopSession
    "end_time": true,    // Set by stopSession
  },
}

func startSession(w http.ResponseWriter, r *http.Request) {
  var session models.Toast
  _, err := models.FillModelFromJSON(r, &session, config)
  if BFE.HandleError(w, err) {
    return
  }

  err = DB.StartToastSession(&session)
  if BFE.HandleError(w, err) {
    return
  }

  RSP.SendObjectResponse(w, http.StatusCreated, session)
}
```
