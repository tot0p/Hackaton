import React, {useEffect, useState} from 'react';
import { useDrag, useDrop } from 'react-dnd';
import DataTableComponent from "../DataTable/DataTable";
import "./loader2.css";
import LineGraph from "../LineGraph/LineGraph";

const DraggableBox = ({ id, text, index, data,graphData,fields,loading,loaded,moveBox, onDelete ,onLoadData}) => {
    const [selectedGraph, setSelectedGraph] = useState('datatable'); // Initial selected graph
    const [,drag] = useDrag({
        type: 'BOX',
        item: { id, index },
    });
    const [, drop] = useDrop({
        accept: 'BOX',
        hover: (draggedItem) => {
            if (draggedItem.index !== index) {
                moveBox(draggedItem.index, index);
                draggedItem.index = index;
            }
        },
    });
    useEffect(() => {
        if (!loaded && !loading){
            onLoadData(id);
        }

    }, []);




    // Define your data visualization components here, you can have multiple
    // graph components and toggle between them using the dropdown.
    const graphComponents = {
        datatable: <DataTableComponent keys={fields} data={data} />,
    };
    if (loaded && graphData !== undefined) {
        if (graphData.length > 0) {
            graphData.forEach((graph) => {
                if (graph.type === "Line") {
                    graphComponents[graph.type] =
                        <LineGraph data={data} xAxisField={graph.xAxis} SeriesField={graph.series}/>;
                }
            });
        }
    }

    const handleChangeGraph = (event) => {
        setSelectedGraph(event.target.value);
    };

    const content = loading ? (
        <div className="loader-5 center"><span></span></div>
    ) : (
        <div>
            <div className="dashboard_controls">
                <select value={selectedGraph} onChange={handleChangeGraph}>
                    {Object.keys(graphComponents).map((graph) => (
                        <option key={graph} value={graph}>
                            {graph}
                        </option>
                    ))}
                </select>
                <button onClick={() => onDelete(id)} className="dashboard_delete-button">
                    X
                </button>
            </div>
            <div className="dashboard_graph">
                {graphComponents[selectedGraph]}
            </div>

        </div>
    );

    return (
        <div ref={(node) => drag(drop(node))} className="dashboard_draggable-box">
            {content}
        </div>
    );
};

export default DraggableBox;
