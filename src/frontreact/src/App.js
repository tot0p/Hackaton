import './App.css';
import Footer from "./components/footer/footer";
import Navbar from "./components/navbar/navbar";
import Dashboard from "./components/Playground/Dashboard";
import {DndProvider} from "react-dnd";
import {HTML5Backend} from "react-dnd-html5-backend";
import MapRenderer from "./components/MapPoints/index";
import Donuttest from './components/donut/donutest';
import Testverticalbarchart from './components/VerticalBarChart/testVerticalBarCharts'; 
import WorldMap from './components/worldMap/worldmap';

function App() {
  return (
    <div className="App">
        <Navbar/>
        <DndProvider backend={HTML5Backend}>
        <Dashboard/>
        </DndProvider>
        <MapRenderer/>
        <Donuttest/>
        <Testverticalbarchart/>
        <WorldMap/>
        <Footer/>
    </div>
  );
}

export default App;
