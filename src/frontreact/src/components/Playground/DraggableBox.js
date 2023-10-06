import React, {useEffect, useState} from 'react';
import { useDrag, useDrop } from 'react-dnd';
import DataTableComponent from "../DataTable/DataTable";
import "./loader2.css";
import LineGraph from "../LineGraph/LineGraph";
import MapMarker from "../MapMaker/MapMarker";
import MapArea from "../mapWithVectorLayers/mapWithVectorLayers";

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

    useEffect(() => {
        window.addEventListener('resize', () => {
            setSelectedGraph('datatable');
        });
    }, []);


    // Define your data visualization components here, you can have multiple
    // graph components and toggle between them using the dropdown.
    const graphComponents = {
        datatable: <DataTableComponent keys={fields} data={data} />,
    };
    if (loaded && graphData !== undefined) {
        if (graphData.length > 0) {
            graphData.forEach((graph) => {
                console.log(graph);
                if (graph.type === "Line") {
                    graphComponents[graph.type] =
                        <LineGraph data={data} xAxisField={graph.xAxis} SeriesField={graph.series}/>;
                }
                if (graph.type === "MapMarker") {
                    // generate markers
                    let markers = [];
                    data.forEach((item) => {
                        // precompute label
                        let Label = ""
                        if (item[graph.position.posName] !== null) {
                        graph.labelsName.forEach((labelName) => {
                            if (item[labelName] !== null){
                                Label += item[labelName].toString() + ", ";
                            }
                        });

                            markers.push({
                                lat: item[graph.position.posName][graph.position.latName],
                                lng: item[graph.position.posName][graph.position.logName],
                                label: Label
                            });
                        }
                    });
                    graphComponents[graph.type] = <MapMarker markers={markers}/>;
                }
                if (graph.type === "MapArea") {
                    // generate areas
                    let areas = [];
                    data.forEach((item) => {
                        if (item[graph.area.areaName] !== null) {
                            // get the area
                            if (item[graph.area.areaName][graph.area.geoName]["type"] === "Polygon") {
                                let area = item[graph.area.areaName][graph.area.geoName][graph.area.cordName][0]
                                console.log("area", area);
                                console.log("id", item["id_eqpt"]);
                                areas.push({
                                    area
                                });
                            }

                        }
                    });
                    console.log(areas);
                    graphComponents[graph.type] = <MapArea areas={areas} areaColor={graph.color}/>;
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
        <div className="dashboard_box">
            <div className="dashboard_title">{text}</div>
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
