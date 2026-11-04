package main
import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"github.com/MobilityData/gtfs-realtime-bindings/golang/gtfs"
	"google.golang.org/protobuf/proto"
)


// getVehicles now sends a list of coordinates
func getVehicles(w http.ResponseWriter, r *http.Request) {
	// sets response headers
	// tells browser you're returning json
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	// download the protobuf feed
	resp, err := http.Get("https://svc.metrotransit.org/mtgtfs/vehiclepositions.pb")
	if err != nil {
		http.Error(w, "failed to fetch", http.StatusInternalServerError)
		return
	}
	// close HTTP response body to prevent resource leak
	defer resp.Body.Close()
	// read response into bytes
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "failed to read response body", http.StatusInternalServerError)
		return
	}
	// decode protobuf bytes into a struct
	feed := &gtfs.FeedMessage{} // makes an empty GTFS feed object
	// fills it with decoded data
	if err := proto.Unmarshal(bodyBytes, feed); err != nil {
		http.Error(w, "failed to decode protobuf", http.StatusInternalServerError)
		return
	}
	// backticks means when converted to JSON, the field name should be "id" instead of "ID"
	type VehiclePoint struct {
		ID string `json:"id"`
		Lat float32 `json:"lat"`
		Lon float32 `json:"lon"`
	}
	results := make([]VehiclePoint, 0)

	for _, entity := range feed.Entity {
		vehicle := entity.Vehicle
		if vehicle == nil || vehicle.Vehicle == nil || vehicle.Position == nil {
			continue
		}

		// pointers in the generated structs, so we dereference them with *
		id := *vehicle.Vehicle.Id
		lat := *vehicle.Position.Latitude
		lon := *vehicle.Position.Longitude

		results = append(results, VehiclePoint{
			ID:  id,
			Lat: lat,
			Lon: lon,
		})
	}

	// write JSON to response
	if err := json.NewEncoder(w).Encode(results); err != nil {
		http.Error(w, "failed to encode json", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/api/vehicles", getVehicles)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "TC Map API is Running")
	})
	fmt.Println("Backend starting on http://0.0.0.0:8080...")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}