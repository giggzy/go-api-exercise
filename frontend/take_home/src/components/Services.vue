<template>
  <div class="services">
    <!-- header section -->
      <div class="ui grid">
        <div class="left floated six wide column">
          <div class="ui container">
            <!-- left floated -->
            <h2 class="ui left aligned header">{{title}}</h2>

            <!-- Search -->
            <div class="ui  search">
              <div class="ui icon input">
                <input 
                  v-on:keyup.enter="doSearch"
                  class="prompt" 
                  type="text" 
                  v-model="searchTerm"
                  placeholder="Search services...">
                <i class="search icon"></i>
              </div>
              <div class="results"></div>
            </div>

          </div>
        </div>
        <div class="right floated six wide column">
          <div class="ui container">
            <!-- right floated -->
            <div class="ui primary right aligned button">
              Create Service
            </div>
          </div>
        </div>
      </div>

    <!-- The Service cards -->
    <div class="ui four cards">
      <div class="card" v-for="service in services" :key="service.id">
          <div class="content">
              <div class="header">{{service.name}}</div>
              <div class="description"> {{service.description}} </div>
              <div class="url">
                  <a href="{{service.url}}">link</a>
              </div>
              <div class="url_text">
                  {{service.url}}
              </div>
          </div>
          <div class="ui bottom attached button">
              <i class="add icon"></i>
              Add Friend
          </div>
        </div> <!-- end card -->
      </div> <!-- end cards -->

      <Pagination/>
        
    </div>
</template>

<script>
import axios from 'axios';
import Pagination from './Pagination.vue';

const BASE_API_URL = 'http://localhost:8080/';
async function makeAPICall (query) {
            console.log("makeAPICall invoked: " + query);
            const response = await axios.get(BASE_API_URL + query);
            return response;
        }

export default {
    props: {
        title: {
            type: String,
            required: true
        }
    },
    async mounted() {
        try {
            console.log(BASE_API_URL + "services");
            const response = await makeAPICall("services");
            this.services = response.data;
        }
        catch (e) {
            console.error(e);
        }
    },
    data() {
        return {
            services: [],
            searchTerm: "",
        };
    },
    methods: {
        onClick: function () {
            this.$emit("clicked");
            console.log("clicked");
        },
        doSearch: async function (event) {
            if ( this.searchTerm.length === 0 ) {
              // Just get all the services
              const response = await makeAPICall("services");
              console.log(response);
              this.services = response.data;
            } else {
              //console.log("search invoked: " + this.searchTerm);
              const response = await makeAPICall("services?search=" + this.searchTerm);
              this.services = response.data;
            }
        },
    },
    components: { Pagination }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>