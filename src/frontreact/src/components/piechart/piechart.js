import React, { useEffect } from 'react';
import * as echarts from 'echarts';

const PieChart = ({ data }) => {
  useEffect(() => {
    const chart = echarts.init(document.getElementById('pie-chart'));

    // Les données du graphique
    const option = {
      title: {
        text: '% de choses données',
        left: 'center',
      },
      tooltip: {
        trigger: 'item',
        formatter: '{a} <br/>{b}: {c} ({d}%)',
      },
      series: [
        {
          name: 'Répartition',
          type: 'pie',
          radius: '50%',
          data,
          emphasis: {
            itemStyle: {
              shadowBlur: 10,
              shadowOffsetX: 0,
              shadowColor: 'rgba(0, 0, 0, 0.5)',
            },
          },
        },
      ],
    };

    // Configuration du graphique
    chart.setOption(option);

    // Nettoyer le graphique lorsque le composant est démonté
    return () => {
      chart.dispose();
    };
  }, [data]);

  return <div id="pie-chart" style={{ width: '100%', height: '400px' }} />;
};

export default PieChart;
