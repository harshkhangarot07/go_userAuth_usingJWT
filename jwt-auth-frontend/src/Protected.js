import React, { useEffect, useState } from 'react';
import axios from 'axios';

const Protected = () => {
  const [message, setMessage] = useState('');

  useEffect(() => {
    const fetchProtectedData = async () => {
      try {
        const token = localStorage.getItem('authToken');
        const response = await axios.get('/protected', {
          headers: { Authorization: `Bearer ${token}` },
        });
        setMessage(response.data);
      } catch (error) {
        console.error('Access error:', error);
        setMessage('Access denied');
      }
    };

    fetchProtectedData();
  }, []);

  return (
    <div>
      <h2>Protected Route</h2>
      <p>{message}</p>
    </div>
  );
};

export default Protected;
