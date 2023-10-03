// DraggableBox.js
import React from 'react';
import { useDrag, useDrop } from 'react-dnd';

const DraggableBox = ({ id, text, index, moveBox, onDelete }) => {
    const [, drag] = useDrag({
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

    return (
        <div ref={(node) => drag(drop(node))} className="dashboard_draggable-box">
            {text}
            <button onClick={() => onDelete(id)} className="dashboard_delete-button">
                X
            </button>
        </div>
    );
};

export default DraggableBox;
