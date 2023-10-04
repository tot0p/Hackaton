import React, {useEffect, useState} from 'react';
import { useDrag, useDrop } from 'react-dnd';
import DataTableComponent from "../DataTable/DataTable";

const DraggableBox = ({ id, text, index, data,fields,loading,loaded,moveBox, onDelete ,onLoadData}) => {
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

    const handleChangeGraph = (event) => {
        setSelectedGraph(event.target.value);
    };

    const content = loading ? (
        <p>Loading data...</p>
    ) : (
        <div>
            <div className="dashboard_controls">
                <select value={selectedGraph} onChange={handleChangeGraph}>
                    <option value="datatable">Data Table</option>
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