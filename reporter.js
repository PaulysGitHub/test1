var reporter = require('cucumber-html-report');

var options = {
    theme: 'bootstrap',
    jsonFile: 'test/report.json',
    output: 'test/report.html',
    reportSuiteAsScenarios: true,
    launchReport: true,
    metadata: {
        "App Version": "0.3.2",
        "Test Environment": "STAGING",
        "Browser": "Chrome 54.0.2840.98",
        "Platform": "Windows 10",
        "Parallel": "Scenarios",
        "Executed": "Remote"
    }
};

reporter.generate(options);
// From terminal run
// node reporter.js to view test results report

// For more indepth reports run both of these commands
// godog --format=cucumber > log/report.json
// after the test runs look at the report
// node reporter.js to view more vivid test results report