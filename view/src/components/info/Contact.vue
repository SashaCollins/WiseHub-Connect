<template>
  <div>
    <div class="container">
        <div class="header">
            <h2>WiseHub-Contact</h2>
        </div>
        <b-container fluid="true" class="container-fluid">
            <div class="container" id="containerContact">
                <div style="text-align:center">
                    <h3>Contact Us</h3>
                    <p>Swing by for a cup of coffee, or leave us a message:</p>
                </div>
                <div class="row">
                    <div class="column">
                        <img src="../../assets/wisehubIcon.png" style="width:100%">
                    </div>
                    <div class="column">
                        <form @submit.prevent="handleSubmit">
                            <label for="fname">First Name</label>
                            <input type="text" class="form-control" id="fname" name="firstname" placeholder="Neville" required>
                            <div class="valid-feedback">Valid.</div>
                            <div class="invalid-feedback">Please fill out this field.</div>

                            <label for="lname">Last Name</label>
                            <input type="text" id="lname" name="lastname" placeholder="Longbottom">

                            <label for="email">Email</label>
                            <input type="email" class="form-control" id="email" name="email" placeholder="nlongbottom@hogwarts.edu" required>

                            <label for="country">Country</label>
                            <select class="custom-select" id="country" name="country">
                                <option selected>Select Location:</option>
                                <option value="germany">Germany</option>
                                <option value="other">Other</option>
                            </select>

                            <label for="subject">Subject</label>
                            <textarea  id="subject" name="subject" placeholder="I can't remember my password..." style="height:170px" required></textarea>
                            <input type="submit" value="Submit">

                        </form>
                    </div>
                </div>
            </div>
        </b-container>
    </div>
  </div>
</template>

<script>
    export default {
        name: "Contact",
        methods: {
            handleSubmit: function () {

            }
        }
    }
</script>

<style scoped lang="scss">

    h3 {
      font-size: 50px;
    }

    * {
        box-sizing: border-box;
    }

    /* Style inputs */
    input[type=text], select, textarea, #email {
        width: 100%;
        border: 1px solid #ccc;
        border-radius: 10px;
        margin-top: 6px;
        margin-bottom: 16px;
        resize: vertical;
        background-color: #001919;
        color: #ccc;
    }
    input[type=text], textarea, #email {
        padding: 12px;
    }

    input[type=submit] {
        background-color: #8b0000;
        color: white;
        padding: 12px 20px;
        border: none;
        border-radius: 10px;
        cursor: pointer;
    }

    input[type=submit]:hover {
        background-color: #9a2020;
    }
    .container-fluid {
      padding-top: 30px;
    }
    /* Style the container/contact section */
    #containerContact {
        width: 75%;
        border-radius: 10px;
        background-color: #002828;
        padding: 10px;
        color: #f2f2f2;
    }

    /* Create two columns that float next to eachother */
    .column {
        float: left;
        width: 50%;
        margin-top: 6px;
        padding: 20px;
    }

    /* Clear floats after the columns */
    .row:after {
        content: "";
        display: table;
        clear: both;
    }

    /* Responsive layout - when the screen is less than 600px wide, make the two columns stack on top of each other instead of next to each other */
    @media screen and (max-width: 650px) {
        .column, input[type=submit] {
            width: 100%;
            margin-top: 0;
        }
    }
    @media screen and (max-width: 975px) {
        .column, input[type=submit] {
            width: 100%;
            margin-top: 0;
        }
    }
</style>