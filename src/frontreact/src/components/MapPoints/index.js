import React from "react";
import MapPoints from "./MapPoints";

const MapRenderer = () => {
  const data = [
    [10, 20], // Example data point [x, y]
    [30, 40],
    [120,200]
  ];

  return (
    <div className="App">
      <h1>2-Axis Map with ECharts</h1>
      <MapPoints data={data} />
    </div>
  );
};

export default MapRenderer;