import L from "leaflet";
import React, { Component } from 'react';
import { MapContainer, TileLayer, Polygon} from 'react-leaflet';
import 'leaflet/dist/leaflet.css'

const MapArea = ({ areas,areaColor }) => {
  return (
      <div>
        <MapContainer
            center={[0,0]}
            zoom={13}
            style={{ width: '100%', height: '400px' }}
        >
          <TileLayer
              url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
              attribution='&copy; OpenStreetMap contributors'
          />

          {/* Boucle à travers les points et ajoute des marqueurs sur la carte */}
          {areas.map(area => (
              <Polygon
                  positions={area["area"].map((item) => {
                    return [item[1],item[0]];
                  })}
                  pathOptions={{
                    color: areaColor,
                    fillColor: areaColor,
                    fillOpacity: 0.4,
                  }
                    }
                />
            ))}
        </MapContainer>
      </div>
  );
}

export default MapArea;