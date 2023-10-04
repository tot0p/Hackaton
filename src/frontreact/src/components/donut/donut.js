import React, { useEffect, useRef } from "react";
import * as echarts from "echarts";

const Donut = ({ data, labels }) => {
  const chartRef = useRef(null);

  useEffect(() => {
    // Créez le graphique en forme de donut
    const chart = echarts.init(chartRef.current);

    // Définissez les options du graphique
    const options = {
      tooltip: {
        trigger: "item",
        formatter: "{a} <br/>{b}: {c} ({d}%)",
      },
      legend: {
        orient: "vertical",
        left: 10,
        data: labels,
      },
      series: [
        {
          name: "Donut Chart",
          type: "pie",
          radius: ["50%", "70%"],
          avoidLabelOverlap: false,
          label: {
            show: false,
            position: "center",
          },
          emphasis: {
            label: {
              show: true,
              fontSize: "20",
              fontWeight: "bold",
            },
          },
          labelLine: {
            show: false,
          },
          data: data.map((value, index) => ({
            value,
            name: labels[index],
          })),
        },
      ],
    };

    // Définissez les options du graphique
    chart.setOption(options);

    // Nettoyez le graphique lors du démontage du composant
    return () => {
      chart.dispose();
    };
  }, [data, labels]);

  return <div ref={chartRef} style={{ width: "100%", height: "400px" }} />;
};

export default Donut;
