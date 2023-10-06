import React, { useEffect, useState } from 'react';
import Modal from 'react-modal';
import axios from "axios";

Modal.setAppElement('#root'); // Set the root element for screen readers

const AddBoxModal = ({ isOpen, onRequestClose, onConfirm }) => {
    const [dataSources, setDataSources] = useState([]);
    const [selectedDataSource, setSelectedDataSource] = useState('');
    const [isLoading, setIsLoading] = useState(false);
    const [searchText, setSearchText] = useState('');

    const fetchData = async () => {
        setIsLoading(true);
        try {
            const sources = await fetchDataSources();
            // sort the sources alphabetically
            sources.sort();
            setDataSources(sources);
            if (sources.length > 0) {
                setSelectedDataSource(sources[0]); // Set the first source as the default
            }
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

    const filteredDataSources = dataSources.filter((source) =>
        source.toLowerCase().includes(searchText.toLowerCase())
    );

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
                    <input
                        type="text"
                        placeholder="Search..."
                        value={searchText}
                        onChange={(e) => setSearchText(e.target.value)}
                    />
                    <select
                        value={selectedDataSource}
                        onChange={(e) => setSelectedDataSource(e.target.value)}
                    >
                        <option value="">-- Select a data source --</option>
                        {filteredDataSources.map((source) => (
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
        const response = await axios.get('http://hackaton.sytes.net:8080/api/data/names');
        return response.data;
    } catch (error) {
        console.error('Error fetching data sources:', error);
        return [];
    }
};

export default AddBoxModal;
