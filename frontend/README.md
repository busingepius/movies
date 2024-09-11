# Movie Rating Frontend
## Getting Started
### Running the React App Locally
1. Open a new terminal in the repo directory and `cd frontend`
- your path (`pwd`) should be **{system directories}/interview-pre-req-check/frontend**
2. `npm install --location=global react-scripts`
3. `npm install`
4. `npm start`
- React app will start at localhost:3000

![img_8.png](../img_8.png)

- Going forward any React call will need to be to the `Forwarded Port` api for port 8080 

![img_9.png](../img_9.png)

- Update the api health endpoint in `frontend/src/Health.js` to the api 8080 `Forwarded Port`

![img_10.png](../img_10.png)

- Validate react app is working properly by seeing it state in terminal it has started and validate it is talking to your api by navigating to `http://localhost:3000/health` and seeing OK in the browser console logs

Troubleshooting:
- If you encounter any issues with `npm install`, try deleting the node_module directory and package-lock.json file in /frontend and running the commands again.
