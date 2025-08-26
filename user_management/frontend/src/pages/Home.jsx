import React, { useState, useEffect } from 'react';
import axios from 'axios';
import {
    ChevronDown,
    ChevronUp,
    Edit,
    Trash,
    Plus,
    X,
    Check,
    Search,
    RefreshCw
} from 'lucide-react';
import './Home.css'
const Home = () => {
    const [users, setUsers] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const [editingUser, setEditingUser] = useState(null);
    const [formData, setFormData] = useState({
        Username: '',
        Password: '',
        Name: '',
        Email: ''
    });
    const [searchTerm, setSearchTerm] = useState('');
    const [sortConfig, setSortConfig] = useState({ key: 'ID', direction: 'ascending' });

    // Fetch all users
    const fetchUsers = async () => {
        setLoading(true);
        try {
            const token = localStorage.getItem("token");

            const response = await axios.get('http://localhost:8080/api/v1/user/getUsers', {
                headers: {
                    Authorization: token
                }
            });
            setUsers(response.data.data);
            setError(null);
        } catch (err) {
            setError('Failed to fetch users. Please try again later.');
            console.error('Error fetching users:', err);
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        fetchUsers();
    }, []);

    // Handle delete user
    const handleDelete = async (username) => {
        if (window.confirm('Are you sure you want to delete this user?')) {
            try {
                const token = localStorage.getItem("token");
                await axios.delete(`http://localhost:8080/api/v1/user/delete/${username}`, {
                    headers: {
                        Authorization: token
                    }
                });
                setUsers(users.filter(user => user.Username !== username));
            } catch (err) {
                setError('Failed to delete user. Please try again.');
                console.error('Error deleting user:', err);
            }
        }
    };

    // Handle starting edit
    const handleEdit = (user) => {
        setEditingUser(user.Username);
        setFormData({
            Id: user.ID,
            Password: '',
            Name: user.Name,
            Email: user.Email
        });
    };

    // Handle form input changes
    const handleInputChange = (e) => {
        const { name, value } = e.target;
        setFormData(prev => ({
            ...prev,
            [name]: value
        }));
    };

    // Handle update submit
    const handleUpdate = async () => {
        try {
            const updatedData = { ...formData, Username: editingUser };
            const token = localStorage.getItem("token");
            await axios.put(`http://localhost:8080/api/v1/user/update/${editingUser}`, updatedData, {
                headers: {
                    Authorization: token
                }
            });
            setUsers(users.map(user =>
                user.Username === editingUser ? { ...user, ...formData } : user
            ));
            setEditingUser(null);
        } catch (err) {
            setError('Failed to update user. Please try again.');
            console.error('Error updating user:', err);
        }
    };

    // Handle sorting
    const requestSort = (key) => {
        let direction = 'ascending';
        if (sortConfig.key === key && sortConfig.direction === 'ascending') {
            direction = 'descending';
        }
        setSortConfig({ key, direction });
    };

    // Get sorted, filtered users
    const getSortedUsers = () => {
        const filtered = users.filter(user =>
            user.Username.toLowerCase().includes(searchTerm.toLowerCase())
        );

        return [...filtered].sort((a, b) => {
            if (a[sortConfig.key] < b[sortConfig.key]) {
                return sortConfig.direction === 'ascending' ? -1 : 1;
            }
            if (a[sortConfig.key] > b[sortConfig.key]) {
                return sortConfig.direction === 'ascending' ? 1 : -1;
            }
            return 0;
        });
    };

    // Get sort indicator
    const getSortIndicator = (key) => {
        if (sortConfig.key === key) {
            return sortConfig.direction === 'ascending' ?
                <ChevronUp className="icon inline-icon" /> :
                <ChevronDown className="icon inline-icon" />;
        }
        return null;
    };

    return (
        <div className="container">
            <div className="card">
                <div className="header">
                    <h1 className="header-title">User Management</h1>
                </div>

                {/* Search bar */}
                <div className="search-container">
                    <div className="search-wrapper">
                        <Search className="search-icon icon" />
                        <input
                            type="text"
                            placeholder="Search users..."
                            className="search-input"
                            value={searchTerm}
                            onChange={e => setSearchTerm(e.target.value)}
                        />
                    </div>
                </div>

                {/* Error message */}
                {error && (
                    <div className="alert alert-error" role="alert">
                        <p>{error}</p>
                    </div>
                )}


                {/* User table */}
                <div className="table-container">
                    <table className="table">
                        <thead>
                            <tr>
                                <th onClick={() => requestSort('ID')}>
                                    ID {getSortIndicator('ID')}
                                </th>
                                <th onClick={() => requestSort('Username')}>
                                    Username {getSortIndicator('Username')}
                                </th>
                                <th onClick={() => requestSort('Name')}>
                                    Name {getSortIndicator('Name')}
                                </th>
                                <th onClick={() => requestSort('Email')}>
                                    Email {getSortIndicator('Email')}
                                </th>
                                <th style={{ textAlign: 'right' }}>Actions</th>
                            </tr>
                        </thead>
                        <tbody>
                            {loading ? (
                                <tr>
                                    <td colSpan="5">
                                        <div className="loading-container">
                                            <div className="spinner"></div>
                                            <span>Loading...</span>
                                        </div>
                                    </td>
                                </tr>
                            ) : getSortedUsers().length === 0 ? (
                                <tr>
                                    <td colSpan="5" className="empty-message">
                                        No users found
                                    </td>
                                </tr>
                            ) : (
                                getSortedUsers().map(user => (
                                    <tr key={user.ID}>
                                        <td className="text-primary">
                                            {user.ID}
                                        </td>
                                        <td className="text-primary">
                                            {user.Username}
                                        </td>
                                        <td className="text-primary">
                                            {editingUser === user.Username ? (
                                                <input
                                                    type="text"
                                                    name="Name"
                                                    value={formData.Name}
                                                    onChange={handleInputChange}
                                                    className="form-input"
                                                />
                                            ) : (
                                                user.Name
                                            )}
                                        </td>
                                        <td className="text-secondary">
                                            {editingUser === user.Username ? (
                                                <input
                                                    type="email"
                                                    name="Email"
                                                    value={formData.Email}
                                                    onChange={handleInputChange}
                                                    className="form-input"
                                                />
                                            ) : (
                                                user.Email
                                            )}
                                        </td>
                                        <td>
                                            {editingUser === user.Username ? (
                                                <div className="table-actions">
                                                    <X
                                                        className="icon-md action-icon cancel-icon"
                                                        onClick={() => setEditingUser(null)}
                                                    />
                                                    <Check
                                                        className="icon-md action-icon confirm-icon"
                                                        onClick={handleUpdate}
                                                    />
                                                </div>
                                            ) : (
                                                <div className="table-actions">
                                                    <Edit
                                                        className="icon-md action-icon edit-icon"
                                                        onClick={() => handleEdit(user)}
                                                    />
                                                    <Trash
                                                        className="icon-md action-icon delete-icon"
                                                        onClick={() => handleDelete(user.Username)}
                                                    />
                                                </div>
                                            )}
                                        </td>
                                    </tr>
                                ))
                            )}
                        </tbody>
                    </table>
                </div>

                <div className="footer">
                    <p className="footer-text">
                        Showing {getSortedUsers().length} users
                    </p>
                </div>
            </div>
        </div>
    );
};

export default Home;