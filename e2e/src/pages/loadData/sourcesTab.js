const loadDataPage = require(__srcdir + '/pages/loadData/loadDataPage.js');
const { By } = require('selenium-webdriver');

const dataWriteMethFilter = '[data-testid=input-field][placeholder*=\'Search\']';
const clientLibsSection = '[data-testid="write-data--section client-libraries"]';
const telegrafPlugSection = '[data-testid="write-data--section telegraf-plugins"]';

class sourcesTab extends loadDataPage {

    constructor(driver) {
        super(driver);
    }

    async isTabLoaded(tabUrlPart, selectors = undefined){
        await super.isTabLoaded(tabUrlPart, [
            {type:'css', selector:dataWriteMethFilter},
            {type:'css', selector:clientLibsSection},
            {type:'css', selector:telegrafPlugSection}
        ]);
    }

}

module.exports = sourcesTab;
