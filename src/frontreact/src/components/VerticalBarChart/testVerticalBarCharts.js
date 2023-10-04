import React from "react";
import VerticalBarChart from "./VerticalBarChart";

const Testverticalbarchart = () => {
    const data = [1, 20, 20, 20]; // Exemple de données
    const labels = ["Le qi de Lulu", "Le qi de Sambre", "Le qi de schnek", "Le qi de Tot0p"]; // Exemple d'étiquettes

  return (
    <div className="App">
      <h1>G la barre</h1>
      <VerticalBarChart data={data} labels={labels} />
    </div>
  );
};

export default Testverticalbarchart;