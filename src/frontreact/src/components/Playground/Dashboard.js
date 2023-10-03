// Dashboard.js
import React, { useState } from 'react';
import DraggableBox from './DraggableBox';
import AddBoxModal from './AddBoxModal';
import './Dashboard.css';


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
                    moveBox={moveBox}
                    onDelete={handleDeleteBox}
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

export default Dashboard;
