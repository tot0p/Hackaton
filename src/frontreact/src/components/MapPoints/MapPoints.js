import React, { useEffect, useRef } from "react";
import * as echarts from "echarts";

const MapPoints = ({ data }) => {
  const chartRef = useRef(null);

  useEffect(() => {
    // Initialize the ECharts instance
    const chart = echarts.init(chartRef.current);

    // Define your chart options
    const options = {
      tooltip: {},
      xAxis: {},
      yAxis: {},
      series: [
        {
          type: "scatter", // Use scatter plot for points
          data: data, // Your data should be an array of points
        },
      ],
    };

    // Set the chart options
    chart.setOption(options);

    // Clean up the chart instance when the component unmounts
    return () => {
      chart.dispose();
    };
  }, [data]);

  return <div ref={chartRef} style={{ width: "100%", height: "500px" }}></div>;
};

export default MapPoints;