import React from 'react';
import PieChart from './piechart';

const Testpiechart= () => {
  const data = [
    { name: 'Puissance de lulu', value: 5 },
    { name: 'Puissance de Sambre', value: 100 },
    { name: 'Puissance de Axel', value: 70 },
    { name: 'Puissance de Tot0p', value: 90 },
  ];

  return (
    <div>
      <h1>% de datas</h1>
      <PieChart data={data} />
    </div>
  );
};

export default Testpiechart;
