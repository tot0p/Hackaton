import React from 'react';
import DataTable from 'react-data-table-component'

const DataTableComponent = ({ keys,data }) => {
    // field is an array of strings that represent the keys of the data
    // sort key by alphabetical order
    keys.sort();
    // Define the columns for the data table
    const columns = keys.map((key) => ({
        name: key,
        selector: key,
        sortable: true,
    }));
    return (
        <DataTable
            columns={columns}
            data={data}
            direction="auto"
            pagination
            fixedHeader
            fixedHeaderScrollHeight="500px"
            highlightOnHover
            responsive
            subHeaderAlign="right"
            subHeaderWrap
        />
    );
};


export default DataTableComponent;