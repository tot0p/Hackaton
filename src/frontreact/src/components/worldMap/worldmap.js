import React, { Component } from 'react';
import { MapContainer, TileLayer, Marker, Popup } from 'react-leaflet';
import 'leaflet/dist/leaflet.css'
import l from 'leaflet';

class WorldMap extends Component {
  constructor(props) {
    super(props);
    this.state = {
      // Exemple de données de points (latitude et longitude)
      points: [
        { id: 1, lat: 51.505, lng: -0.09, label: 'Point 1' },
      ],
    };
  }

  render() {
    const { points } = this.state;

    return (
      <div>
        <h1>Carte du Monde</h1>
        <MapContainer
          center={[0, 0]}
          zoom={2}
          style={{ width: '100%', height: '400px' }}
        >
          <TileLayer
            url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
            attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
          />

          {/* Boucle à travers les points et ajoute des marqueurs sur la carte */}
          {points.map(point => (
            <Marker
              key={point.id}
              position={[point.lat, point.lng]}
            >
              <Popup>{point.label}</Popup>
            </Marker>
          ))}
        </MapContainer>
      </div>
    );
  }
}

export default WorldMap;
