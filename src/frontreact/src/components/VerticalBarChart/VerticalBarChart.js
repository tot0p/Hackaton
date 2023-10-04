import React, { useEffect, useRef } from "react";
import * as echarts from "echarts";

const VerticalBarChart = ({ data, labels }) => {
  const chartRef = useRef(null);

  useEffect(() => {
    // Créez le graphique à barres verticales
    const chart = echarts.init(chartRef.current);

    // Définissez les couleurs personnalisées pour chaque barre
    const customColors = ["#FF5733", "#FFC300", "#C70039", "#900C3F"];

    // Définissez la couleur du texte de la légende
    const legendTextColor = "#333"; // Couleur personnalisée

    // Définissez la couleur du petit bouton de la légende
    const legendItemColor = "#FF5733"; // Couleur personnalisée

    // Définissez les options du graphique
    const options = {
      tooltip: {
        trigger: "axis",
        axisPointer: {
          type: "shadow",
        },
      },
      legend: {
        data: ["Valeur"],
        textStyle: {
          color: legendTextColor,
        },
        itemStyle: {
          color: legendItemColor,
        },
      },
      xAxis: {
        data: labels,
        axisLabel: {
          interval: 0,
          rotate: 45,
        },
      },
      yAxis: {},
      series: [
        {
          name: "Valeur",
          type: "bar",
          data: data,
          barWidth: "60%",
          itemStyle: {
            // Utilisez le tableau customColors pour définir les couleurs personnalisées
            color: (params) => customColors[params.dataIndex],
          },
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

export default VerticalBarChart;
