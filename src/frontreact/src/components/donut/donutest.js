import React from "react";
import DonutChart from "./donut";

const Donuttest = () => {
  const data = [1, 20, 20, 20]; // Exemple de données
  const labels = ["Le qi de Lulu", "Le qi de Sambre", "Le qi de schnek", "Le qi de Tot0p"]; // Exemple d'étiquettes

  return (
    <div className="App">
      <h1>Sa mère le donut</h1>
      <DonutChart data={data} labels={labels} />
    </div>
  );
};

export default Donuttest;