# tc_map
## A live map of Metro Transit vehicles in the Twin Cities.
1. How it works: the Go backend downloads Metro Transit’s GTFS-realtime feed (protobuf), decodes it, and serves /api/vehicles as JSON.
    - React + Mapbox polls that endpoint every 10 seconds and draws the buses.
2. How to run it: create a .env with REACT_APP_MAPBOX_ACCESS_TOKEN=..., then run docker compose up, then open localhost:3000.
3. Next: click a bus for details, WebSockets instead of polling, Redis cache. Move your TODO comments from App.js here.