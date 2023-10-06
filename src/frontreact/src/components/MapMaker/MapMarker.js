import L from "leaflet";
import React, { Component } from 'react';
import { MapContainer, TileLayer, Marker, Popup } from 'react-leaflet';
import 'leaflet/dist/leaflet.css'

const MapMarker = ({ markers }) => {
  return (
    <div>
      <MapContainer
        center={markers[0]}
        zoom={13}
        style={{ width: '100%', height: '400px' }}
      >
        <TileLayer
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
          attribution='&copy; OpenStreetMap contributors'
        />

        {/* Boucle à travers les points et ajoute des marqueurs sur la carte */}
        {markers.map(marker => (
          <Marker
            position={[marker.lat, marker.lng]}
            icon={new L.Icon({
                iconUrl: 'https://cdn0.iconfinder.com/data/icons/small-n-flat/24/678111-map-marker-512.png',
                iconSize: [25, 25],
                iconAnchor: [12, 25],
                popupAnchor: [1, -34],
                }
            )}
          >
            <Popup>{marker.label}</Popup>
          </Marker>
        ))}
      </MapContainer>
    </div>
  );
}

export default MapMarker;