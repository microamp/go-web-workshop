/*
 * Copygright 2016 Google Inc. All rights reserved.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to writing, software distributed
 * under the License is distributed on a "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied.
 *
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

function EventsCtrl($scope, $http) {
  $scope.events = [];
  $scope.newEvent = {};
  $scope.working = false;

  var logError = function(data, status) {
    alert('code '+status+': '+data);
    $scope.working = false;
  };

  var refresh = function() {
    return $http.get('/api/events').
      success(function(data) { $scope.events = data; }).
      error(logError);
  };

  $scope.addEvent = function() {
    $scope.working = true;
    $http.post('/api/events', $scope.newEvent).
      error(logError).
      success(function(data) {
        refresh().then(function() {
          $scope.working = false;
          $scope.newEvent = {};
          setTimeout(refresh, 1000);
        });
      });
  };

  refresh().then(function() { $scope.working = false; });
}