async function Robot() {
    const data = await fetch('http://localhost:8080/query', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            query: `
            query {
              videos {
                id
                title
                url
              }
            }
            `
        })
    })
    .then(res => res.json()) 
    .then(data => {
        console.log(data.data.videos); 
        return data; 
    })
    .catch(err => {
        console.error('Error:', err);
    });

    return data;
}

Robot().then(data => console.log('Final Data:', data));