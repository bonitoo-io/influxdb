const { expect, assert } = require('chai');
const { By, until } = require('selenium-webdriver');

const baseSteps = require(__srcdir + '/steps/baseSteps.js');
const sourcesTab = require(__srcdir + '/pages/loadData/sourcesTab.js');

class sourcesSteps extends baseSteps {

    constructor(driver) {
        super(driver);
        this.sourcesTab = new sourcesTab(driver);
    }

    async isLoaded() {
        await this.sourcesTab.isTabLoaded();
    }

}

module.exports = sourcesSteps;


