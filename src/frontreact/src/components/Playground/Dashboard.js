// Dashboard.js
import React, { useState } from 'react';
import DraggableBox from './DraggableBox';
import AddBoxModal from './AddBoxModal';
import './Dashboard.css';
import axios from "axios";


const Dashboard = () => {
    const [boxes, setBoxes] = useState([]);
    const [boxCount, setBoxCount] = useState(0);
    const [isModalOpen, setIsModalOpen] = useState(false);

    const handleAddBox = () => {
        setIsModalOpen(true);
    };

    const handleConfirmAddBox = (selectedDataSource) => {
        setIsModalOpen(false);
        setBoxes([
            ...boxes,
            {
                id: boxCount,
                text: selectedDataSource,
                data: [],
                fields: [],
                loading: false,
                loaded: false,
            },
        ]);
        setBoxCount(boxCount + 1);
    };

    const handleCancelAddBox = () => {
        setIsModalOpen(false);
    };

    const handleDeleteBox = (id) => {
        const updatedBoxes = boxes.filter((box) => box.id !== id);
        setBoxes(updatedBoxes);
    };

    const handleLoadBoxData = (id) => {
        // fetch data for the box with the given id
        const box = boxes.find((box) => box.id === id);
        const DataSources = box.text;
        box.loading = true;
        console.log("DataSources: " + DataSources);
        // update the boxes state
        setBoxes([...boxes]);
        fetchData(DataSources).then((data) => {
            // get the box with the given id
            const box = boxes.find((box) => box.id === id);
            // update the box with the data
            box.data = data;
            fetchDataFields(DataSources).then((fields) => {
                // get the box with the given id
                const box = boxes.find((box) => box.id === id);
                // update the box with the data fields
                box.fields = fields;
                box.loading = false;
                box.loaded = true;
                // update the boxes state
                console.log(boxes);
                updateBoxData();
                console.log("End of " + DataSources);
                console.log(boxes);
            });
        });
    }

    const updateBoxData = () => {
        setBoxes([...boxes])
    }

    const moveBox = (fromIndex, toIndex) => {
        const updatedBoxes = [...boxes];
        const [movedBox] = updatedBoxes.splice(fromIndex, 1);
        updatedBoxes.splice(toIndex, 0, movedBox);
        setBoxes(updatedBoxes);
    };

    return (
        <div className="dashboard">
            <button onClick={handleAddBox} className="dashboard_button">
                Add Box
            </button>
            <div className="dashboard_box-grid">
            {boxes.map((box, index) => (
                <DraggableBox
                    key={box.id}
                    id={box.id}
                    text={box.text}
                    index={index}
                    data={box.data}
                    fields={box.fields}
                    loading={box.loading}
                    loaded={box.loaded}
                    moveBox={moveBox}
                    onDelete={handleDeleteBox}
                    onLoadData={handleLoadBoxData}
                />
            ))}
            </div>
            <AddBoxModal
                isOpen={isModalOpen}
                onRequestClose={handleCancelAddBox}
                onConfirm={handleConfirmAddBox}
            />
        </div>
    );
};

const fetchDataFields = async (name) => {
    try {
        const response = await axios.get("http://localhost:8080/api/data/name/" + name+"/fields?rmId=false&rmMulFields=false");
        return response.data;
    } catch (error) {
        console.error("Error fetching data sources:", error);
        return [];
    }
}
const fetchData = async (name) => {
    try {
        const response = await axios.get("http://localhost:8080/api/data/name/" + name);
        return response.data;
    } catch (error) {
        console.error("Error fetching data sources:", error);
        return [];
    }
}

export default Dashboard;
