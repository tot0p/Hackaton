import React, { useEffect } from 'react';
import * as echarts from 'echarts';

const LineGraph = ({ data,xAxisField,SeriesField }) => {
useEffect(() => {
    const chart = echarts.init(document.getElementById('line-graph'));
    let xAxis = data.map((item) => {
        return item[xAxisField];
    });
    xAxis= xAxis.sort();

    let series = SeriesField.map((SeriesName) => {
        return {
            name: SeriesName,
            type: 'line',
            data: data.map((item) => {
                if (item[SeriesName] === null) {
                    return 0;
                }
                return item[SeriesName];
            }),
            emphasis: {
                itemStyle: {
                    shadowBlur: 10,
                    shadowOffsetX: 0,
                    shadowColor: 'rgba(0, 0, 0, 0.5)',
                }
            }
        }
    });

    // Les données du graphique
    const option = {
      tooltip: {
        trigger: 'item',
        formatter: '{a} <br/>{b}: {c}',
      },
      xAxis: {
        type: 'category',
        data: xAxis,
      },
      yAxis: {
        type: 'value',
      },
      series: series,
        legend: {
            data: SeriesField,
            bottom: 0
        }
    };

    // Configuration du graphique
    chart.setOption(option);

    // Nettoyer le graphique lorsque le composant est démonté
    return () => {
      chart.dispose();
    };
  }, []);

  return <div id="line-graph" style={{ width: '100%', height: '400px' }} />;

}


export default LineGraph;