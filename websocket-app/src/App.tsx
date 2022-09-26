import React from 'react';
 
<ReactPolling
  url={'url to poll'}
  interval= {3000} // in milliseconds(ms)
  retryCount={3} // this is optional
  onSuccess={() => console.log('handle success')}
  onFailure={() => console.log('handle failure')} // this is optional
  method={'GET'}
  headers={headers object} // this is optional
  body={JSON.stringify(data)} // data to send in a post call. Should be stringified always
  render={({ startPolling, stopPolling, isPolling }) => {
    if(isPolling) {
      return (
        <div> Hello I am polling</div>
      );
    } else {
      return (
        <div> Hello I stopped polling</div>
      );
    }
  }}
/>