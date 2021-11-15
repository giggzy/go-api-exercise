<template>
    

<div class="ui cards">
  <div class="card" :v-for="service in services">
    <div class="content">
      <div class="header">service.name</div>
      <div class="description">
      service.description
      </div>
    </div>
    <div class="ui bottom attached button">
      <i class="add icon"></i>
      Add Friend
    </div>
  </div>
</div>

</template>

<script>
export default {
  name: 'Services',
  props: {
    msg: String
  },
  mounted: function () {
      window.axios.get('http://localhost:8080/services')
        .then(response => console.log(response.data))
        .catch(error => {
          console.log(error);
        });
        //.then(response => this.services = response.data)
      // read from api    
      /*
      window.axios.get('https://jsonplaceholder.typicode.com/posts')
        .then(response => this.posts = response.data)
        .catch(error => {
          console.log(error);
        });
        */
    },

    data: function () {
      return {
        services: [ 
            {
                "id": "584548",
                "name": "GABE",
                "description": "b blah, blah, blah",
                "versionCount": 222,
                "url": "https://example.com/b_Service",
                "versions": null,
            },
        ],
      }
    },

  methods: {
    onClick: function () {
      this.$emit('clicked');
      console.log('clicked');
    },

  }
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