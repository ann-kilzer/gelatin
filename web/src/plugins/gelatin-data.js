import axios from 'axios';

const devURL = 'http://localhost:1323';

async function getHelper(url, defaultResponse = null) {
  try {
    const response = await axios.get(url);
    return response.data;
  } catch (error) {
    console.error(error);
    return defaultResponse;
  }
}

export default class MNData {
  constructor(baseURL = devURL) {
    this.baseURL = baseURL;
  }

  async GetLocations() {
    return getHelper(`${this.baseURL}/locations`, []);
  }

  async GetLocation(id) {
    return getHelper(`${this.baseURL}/locations/${id}`);
  }

  async CreateLocation(location) {
    const url = `${this.baseURL}/locations`;
    try {
      await axios.post(url, location);
      return true;
    } catch (error) {
      console.error(error);
      return false;
    }
  }

  async UpdateLocation(location) {
    const url = `${this.baseURL}/locations/${location.id}`;
    try {
      await axios.put(url, location);
      return true;
    } catch (error) {
      console.error(error);
      return false;
    }
  }

  async DeleteLocation(location) {
    const url = `${this.baseURL}/locations/${location.id}`;
    try {
      await axios.delete(url);
      return true;
    } catch (error) {
      console.error(error);
      return false;
    }
  }
}
