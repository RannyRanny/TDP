import axios from 'axios';

const BASE_URL = 'http://localhost:8080'; // Replace with your actual API base URL

export const fetchTasksForUser = async (userId: number): Promise<any[]> => {
    try {
        const response = await axios.get(`${BASE_URL}/tasks/user/${userId}`);
        return response.data;
    } catch (error) {
        console.error("Error fetching tasks:", error);
        return [];
    }
};
