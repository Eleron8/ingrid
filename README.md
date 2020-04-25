# ingrid

go get github.com/Eleron8/ingrid

This small service can show a list of routes between source and each destination

to run service use in route directory `go run main.go` 

Go to http://localhost:8080 
Here is just small message about this service

for taking list of routes go http://localhost:8080/routes with params `src` and `dst`

for example:

http://localhost:8080/routes?src=13.388860,52.517037&dst=13.397634,52.529407&dst=13.428555,52.523219

params:

src - source. It's route's start

dst - destination. It's route's end

# Google App Engine

This service now is hosting on Google App Engine

main page: https://route-service-275310.ew.r.appspot.com

routes: https://route-service-275310.ew.r.appspot.com/routes?src=13.388860,52.517037&dst=13.397634,52.529407&dst=13.428555,52.523219


