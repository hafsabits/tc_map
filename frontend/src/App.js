import React, { useEffect, useRef } from 'react';
import mapboxgl from 'mapbox-gl';
import 'mapbox-gl/dist/mapbox-gl.css';

mapboxgl.accessToken = process.env.REACT_APP_MAPBOX_ACCESS_TOKEN;

function App() {
  const mapContainerRef = useRef(null);
  const mapRef = useRef(null);

  useEffect(() => {
    // 1. Initialize the Map
    mapRef.current = new mapboxgl.Map({
      container: mapContainerRef.current,
      style: 'mapbox://styles/mapbox/streets-v12',
      center: [-93.2323, 44.9723], // UMN HSEC Coordinates
      zoom: 12,
    });

    mapRef.current.addControl(new mapboxgl.NavigationControl(), 'top-right');

    // 2. Wait for map to load before adding the "Source" (the data bucket)
    mapRef.current.on('load', () => {
      mapRef.current.addSource('buses', {
        type: 'geojson',
        data: { type: 'FeatureCollection', features: [] }
      });

      mapRef.current.addLayer({
        id: 'bus-layer',
        type: 'circle',
        source: 'buses',
        paint: {
          'circle-radius': 6,
          'circle-color': '#ff4b5c',
          'circle-stroke-width': 1,
          'circle-stroke-color': '#ffffff'
        }
      });

      // Now that the map is ready, start fetching
      fetchData();
    });

    const fetchData = async () => {
      try {
        const response = await fetch('http://localhost:8080/api/vehicles');
        const rawData = await response.json();

        // Transform Go JSON into Mapbox GeoJSON
        const geojson = {
          type: 'FeatureCollection',
          features: rawData.map(v => ({
            type: 'Feature',
            geometry: {
              type: 'Point',
              coordinates: [v.lon, v.lat]
            },
            properties: { id: v.id }
          }))
        };

        // Update the map source on the fly
        if (mapRef.current && mapRef.current.getSource('buses')) {
          mapRef.current.getSource('buses').setData(geojson);
          console.log(`Updated ${rawData.length} vehicles on map`);
        }
      } catch (err) {
        console.error("Map fetch error:", err);
      }
    };

    const interval = setInterval(fetchData, 10000);

    return () => {
      clearInterval(interval);
      if (mapRef.current) mapRef.current.remove();
    };
  }, []);

  // width and height must be 100% or 100vw/vh for the map to appear
  return (
    <div 
      ref={mapContainerRef} 
      style={{ width: '100vw', height: '100vh', position: 'absolute', top: 0, left: 0 }} 
    />
  );
}

export default App;