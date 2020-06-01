class SidebarService {

    update(selected) {
        localStorage.setItem('theme', selected);
    }
}

export default new SidebarService();