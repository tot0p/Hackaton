// AddBoxModal.js
import React, { useEffect, useState } from 'react';
import Modal from 'react-modal';
import axios from "axios";

Modal.setAppElement('#root'); // Set the root element for screen readers

const AddBoxModal = ({ isOpen, onRequestClose, onConfirm }) => {
    const [dataSources, setDataSources] = useState([]);
    const [selectedDataSource, setSelectedDataSource] = useState('');
    const [isLoading, setIsLoading] = useState(false);

    const fetchData = async () => {
        setIsLoading(true);
        try {
            const sources = await fetchDataSources();
            setDataSources(sources);
            if (sources.length > 0) {
                setSelectedDataSource(sources[0]); // Set the first source as the default
            }
            // add a fake delay to show the loading indicator
        } catch (error) {
            console.error('Error fetching data sources:', error);
        } finally {
            setIsLoading(false);
        }
    };

    useEffect(() => {
        if (isOpen) {
            fetchData();
        }
    }, [isOpen]);

    const handleConfirmAddBox = () => {
        onRequestClose();
        onConfirm(selectedDataSource);
    };

    return (
        <Modal
            isOpen={isOpen}
            onRequestClose={onRequestClose}
            contentLabel="Add Box Modal"
        >
            <h2>Confirm Adding a Box</h2>
            {isLoading ? (
                <p>Loading data sources...</p>
            ) : (
                <>
                    <p>Select a data source for the box:</p>
                    <select
                        value={selectedDataSource}
                        onChange={(e) => setSelectedDataSource(e.target.value)}
                    >
                        {dataSources.map((source) => (
                            <option key={source} value={source}>
                                {source}
                            </option>
                        ))}
                    </select>
                    <button onClick={handleConfirmAddBox}>Add Box</button>
                    <button onClick={onRequestClose}>Cancel</button>
                </>
            )}
        </Modal>
    );
};


export const fetchDataSources = async () => {
    try {
        const response = await axios.get('http://localhost:8080/api/data/names');
        return response.data;
    } catch (error) {
        console.error('Error fetching data sources:', error);
        return [];
    }
};

export default AddBoxModal;
