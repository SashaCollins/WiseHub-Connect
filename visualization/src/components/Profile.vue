<template>
    <div class="container">
        <div>
            Select theme:
            <select v-model="selectedTheme" @change="updateTheme">
                <option
                        v-for="(theme, index) in themes"
                        :key="index"
                        :value="theme.input"
                >
                    {{ theme.name }}
                </option>
            </select>
        </div>
        <hr style="margin: 50px ;border: 1px solid #e3e3e3;">
    </div>
</template>

<script>
    import EventBus from '../bus/event.bus';

    export default {
        name: "Profile",
        data() {
            return {
                themes: [
                    {
                        name: 'WiseHub theme',
                        input: 'wisehub-theme'
                    },
                    {
                        name: 'Black theme',
                        input: 'black-theme'
                    }
                ],
                selectedTheme: 'wisehub-theme',

            }
        },
        methods: {
            updateTheme() {
                this.$store.dispatch('sidebar/theme', this.selectedTheme);
                // localStorage.setItem('theme', this.selectedTheme)
                EventBus.$emit('update_theme', this.selectedTheme);
            }
        }
    }
</script>

<style scoped lang="scss">
    pre {
        font-family: Consolas, monospace;
        color: #000;
        background: #fff;
        border-radius: 2px;
        padding: 15px;
        line-height: 1.5;
        overflow: auto;
    }
</style>