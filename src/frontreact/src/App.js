import './App.css';
import Footer from "./components/footer/footer";
import Navbar from "./components/navbar/navbar";
import Dashboard from "./components/Playground/Dashboard";
import {DndProvider} from "react-dnd";
import {HTML5Backend} from "react-dnd-html5-backend";
import Testpiechart from './components/piechart/testpiechart';
import MapWithVectorLayers from './components/mapWithVectorLayers/mapWithVectorLayers';

function App() {
  return (
    <div className="App">
        <Navbar/> 
        <DndProvider backend={HTML5Backend}>
        <Dashboard/>
        </DndProvider>
        <Footer/>
        <Testpiechart/>
        <MapWithVectorLayers/>
    </div>
  );
}

export default App;
